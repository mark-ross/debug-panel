package debug_panel

import (
	"html/template"
	"log"
	"net/http"
)

type panelManager struct {
	ProgressFlags []*ProgressFlag
	tmpl          *template.Template
}

func (p *panelManager) addProgressFlag(pf *ProgressFlag) {
	p.ProgressFlags = append(p.ProgressFlags, pf)
}

var (
	panel *panelManager
)

func init() {
	panel = &panelManager{
		ProgressFlags: make([]*ProgressFlag, 0, 10),
	}
}

func (p *panelManager) GetHandlerFunc() func(http.ResponseWriter, *http.Request) {
	// load in the template if we haven't already
	if p.tmpl == nil {
		tmpl, err := template.ParseFiles("templates/status-table.gohtml")
		if err != nil {
			panic(err)
		}
		p.tmpl = tmpl
	}

	log.Println("Req received")

	return func(w http.ResponseWriter, r *http.Request) {
		err := p.tmpl.Execute(w, p)
		if err != nil {
			log.Println("Problem executing the template: ", err)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Println("Unable to write anything to the page...")
				panic(err)
			}

		}
	}
}