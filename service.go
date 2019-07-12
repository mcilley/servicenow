package servicenow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"log"
	"os"
	"strings"
	"sort"
)

var httpRE = regexp.MustCompile("^https?://")

// Err represents a possible error message that came back from the server
type Err struct {
	Err    string `json:"error"`
	Reason string `json:"reason"`
}

func (e Err) Error() string {
	if e.Reason == "" {
		return e.Err
	}
	return fmt.Sprintf("%s: %s", e.Err, e.Reason)
}

// Client helps users interact with a ServiceNow instance.
type Client struct {
	Username, Password, Instance string
}

func sys(param string) string {
	return fmt.Sprintf("sysparm_%s", param)
}

// Helper function, for encoding net/url value type into a query for ServiceNow
// Stock net/url encode will join multiple key/value pairs with '&'
// Service now uses "^" for "anding" multiple query params
// See: docs.servicenow.com/bundle/kingston-platform-user-interface/page/use/common-ui-elements/reference/r_OpAvailableFiltersQueries.html
func SnowQueryEncode( v url.Values ) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := k
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('^')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

// PerformFor creates and executes an authenticated HTTP request to an instance,
// for the given table, action and optional id, with the passed options, and
// unmarshal's the JSON into the passed output interface pointer, returning an
// error.
func (c *Client) PerformFor(table, action, id string, opts url.Values, body interface{}, out interface{}) error {
	inst := c.Instance

	if !httpRE.MatchString(inst) {
		inst = "https://" + inst
	}

	u := fmt.Sprintf("%s/%s.do", inst, table)

	vals := url.Values{
		"JSONv2":      {""},
		sys("action"): {action},
	}

	if id != "" {
		vals.Set(sys("sys_id"), id)
		vals.Set("displayvalue", "true")
	}

	if opts != nil {
		vals.Set(sys("query"), SnowQueryEncode(opts))
	}

	meth := http.MethodGet
	buf := &bytes.Buffer{}

	if body != nil {
		meth = http.MethodPost
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(meth, u+"?"+vals.Encode(), buf)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.SetBasicAuth(c.Username, c.Password)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf.Reset()

	// Store JSON so we can do a preliminary error check
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err == nil && echeck.Err != "" {
		return echeck
	}

	return json.NewDecoder(buf).Decode(out)
}

// GetFor performs a servicenow get to the specified table, with options, and
// unmarshal's JSON into the output parameter.
func (c Client) GetFor(table string, id string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "get", id, opts, nil, out)
}

// GetRecordsFor performs a servicenow getRecords to the specified table, with
// options, and unmarshal's JSON into the output parameter.
func (c Client) GetRecordsFor(table string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "getRecords", "", opts, nil, out)
}

// GetRecords performs a servicenow getRecords to the specified table, with
// options, and returns a map for each item
func (c Client) GetRecords(table string, opts url.Values) ([]map[string]interface{}, error) {
	var out struct {
		Records []map[string]interface{}
	}
	err := c.GetRecordsFor(table, opts, &out)
	return out.Records, err
}

// Insert creates a new record for the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) Insert(table string, obj, out interface{}) error {
	return c.PerformFor(table, "insert", "", nil, obj, out)
}

// Update creates a new record for the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) Update(table string, opts url.Values, obj, out interface{}) error {
	return c.PerformFor(table, "update", "", opts, obj, out)
}

// GetFor performs a servicenow get to the specified table, with options, and
// unmarshal's JSON into the output parameter.
func (c Client) DeleteRecord(table string, id string, opts url.Values, obj, out interface{}) error {
	return c.PerformFor(table, "deleteRecord", id , nil, obj, out)
}

// DeleteRecords deletes a record in the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) DeleteRecords(table string, id string, obj, out interface{}) error {
	return c.PerformFor(table, "deleteMultiple", id, nil, obj, out)
}
