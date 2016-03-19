package app

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"regexp"
	"strings"
	"log"
	"io/ioutil"
	. "flotilla"
)

func mustRead(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("mustRead: failed to read file:", filename)
	}
	return b
}

var key_re = regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)

var upload = string(mustRead("upload.html"))

func init() {
  Handle("/").Put(put).Get(get).Options(options)
}

func options(r *http.Request) {
  Status(StatusOK)
}

func put(r *http.Request) {
  c := appengine.NewContext(r)
  key := strings.TrimPrefix(r.URL.Path, "/")
  Ensure(key_re.MatchString(key), StatusForbidden)
  Ensure(r.ContentLength >= 0, StatusLengthRequired)
  Ensure(r.ContentLength <= 128, StatusRequestEntityTooLarge)
  buffer, e := ReadContent(r)
  Check(e)
  value := string(buffer)
  var stored *string
  Check(datastore.RunInTransaction(c, func(c appengine.Context) error {
    v, e := getValue(c, key)
    stored = v
    Check(e)
    if stored == nil {
      Check(putValue(c, key, value))
    }
    return nil
  }, nil))

  if stored == nil {
    Text(StatusCreated, value)
  } else if *stored == value {
    Text(StatusOK, value)
  } else {
    Status(StatusForbidden)
  }
}

func get(r *http.Request) {
  c := appengine.NewContext(r)
  key := strings.TrimPrefix(r.URL.Path, "/")
  Ensure(key_re.MatchString(key), StatusForbidden)
  value, e := getValue(c, key)
  Check(e)
	if value == nil {
		Body(StatusOK, upload, "text/html")
	} else {
		Body(StatusOK, *value, "text/plain; charset=utf-8")
	}
}
