package sdtv
import (
    //"time"
    "regexp"
    "log"
    "github.com/sloonz/go-iconv"
)

type Detail struct {
    Title	    string
    Description	    string
    Date	    string
    WMV		    string
}

func cherr(err error) {
    if err != nil {
	panic(err)
	log.Printf("error: %s\n", err)
    }
}

func (d *Detail) Parse(raw []byte) *Detail{
    re, err := regexp.Compile("http://news.*WMV")
    dumpError(err)

    d.WMV = string(re.Find(raw))
    if d.WMV == "" {
	log.Printf("WMV not found")
    }

    re, err = regexp.Compile("title_3.*>(.*?)<")
    dumpError(err)
    d.Title = string(re.Find(raw))
    if d.Title == "" {
	log.Printf("Title is not found.")
    }

    re, err = regexp.Compile("location2.*black_font.*<span>")
    dumpError(err)
    d.Description = string(re.Find(raw))
    if d.Description == "" {
	log.Printf("Description not found.")
    }

    re, err = regexp.Compile("class=\"mar\">(.*?)<\\/span>")
    dumpError(err)
    d.Date = string(re.Find(raw))
    if d.Date == "" {
	log.Printf("Date not found.")
    }

    return d;
}
