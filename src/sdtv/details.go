package sdtv
import (
    //"time"
    "regexp"
    "log"
)

type Detail struct {
    Title	    string
    Description	    string
    Date	    string
    WMV		    string
}

func (d *Detail) Parse(raw []byte) *Detail{
    re, err := regexp.Compile("http://news.*WMV")
    if err != nil {
	log.Fatal("WMV error")
    }
    d.WMV = string(re.Find(raw))
    if d.WMV == "" {
	log.Printf("WMV not found")
    }

    re, err = regexp.Compile("title_3.*(.*?)")
    if err != nil {
	log.Fatal("Title error")
    }
    d.Title = string(re.Find(raw))
    if d.Title == "" {
	log.Printf("Title not found.")
    }

    re, err = regexp.Compile("新聞摘要")
    if err != nil {
	log.Fatal("Description error")
    }
    d.Description = string(re.Find(raw))
    if d.Description == "" {
	log.Printf("Description not found.")
    }

    re, err = regexp.Compile("class=\"mar\">(.*?)<\\/span>")
    if err != nil {
	log.Fatal("Date error")
    }
    d.Date = string(re.Find(raw))
    if d.Date == "" {
	log.Printf("Date not found.")
    }

    return d;
}
