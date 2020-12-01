package web

/*
 credits due https://gist.github.com/logrusorgru/abd846adb521a6fb39c7405f32fec0cf
*/

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo"
)

//
// Copyright (c) 2018 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See below for more details.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

// ************************************************************************** //
//                                                                            //
// This gist shows a convenient way to load and use HTML templates in Golang  //
// web-applications. The way includes:                                        //
//                                                                            //
//     - ability to reload templates on changes (for development)             //
//     - recursive loading, unlike (html/template).ParseGlob does it          //
//     - short (Rails-like) template name without shared prefix (dir)         //
//       and without extension                                                //
//                                                                            //
// For example, there is a tree of templates                                  //
//         views/                                                             //
//           static/                                                          //
//             home.html                                                      //
//             about.html                                                     //
//             privacypolicy.html                                             //
//             help.html                                                      //
//           user/                                                            //
//             new.html                                                       //
//             edit.html                                                      //
//             show.html                                                      //
//             form.html                                                      //
//           layout/                                                          //
//             head.html                                                      //
//             foot.html                                                      //
//                                                                            //
// Thus, the home.html can include head.html and foor.html following way      //
//                                                                            //
//     {{ template "layout/head" . }}                                         //
//                                                                            //
//     <h1> Home page </h1>                                                   //
//     <!-- other content of the home.html                                    //
//                                                                            //
//     {{ template "layout/foot" . }}                                         //
//                                                                            //
// This is acceptable for user/new and user/edit which can include user/form  //
// along with the layout/head and layout/foot.                                //
//                                                                            //
// ************************************************************************** //

// A Tmpl implements keeper, loader and reloader for HTML templates
type Tmpl struct {
	*template.Template                  // root template
	dir                string           // root directory
	ext                string           // extension
	devel              bool             // reload every time
	funcs              template.FuncMap // functions
	loadedAt           time.Time        // loaded at (last loading time)
}

// NewTmpl creates new Tmpl and loads templates. The dir argument is
// directory to load templates from. The ext argument is extension of
// tempaltes. The devel (if true) turns the Tmpl to reload templates
// every Render if there is a change in the dir.
func NewTmpl(dir, ext string, devel bool, funcs template.FuncMap) (tmpl *Tmpl, err error) {

	// get absolute path
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	tmpl = new(Tmpl)
	tmpl.dir = dir
	tmpl.ext = ext
	tmpl.devel = devel
	tmpl.funcs = funcs

	if err = tmpl.Load(); err != nil {
		tmpl = nil // drop for GC
	}

	return
}

// Dir returns absolute path to directory with views
func (t *Tmpl) Dir() string {
	return t.dir
}

// Ext returns extension of views
func (t *Tmpl) Ext() string {
	return t.ext
}

// Devel returns development pin
func (t *Tmpl) Devel() bool {
	return t.devel
}

// Funcs sets template functions
func (t *Tmpl) Funcs(funcMap template.FuncMap) {
	t.Template = t.Template.Funcs(funcMap)
	t.funcs = funcMap
}

// Load or reload templates
func (t *Tmpl) Load() (err error) {

	// time point
	t.loadedAt = time.Now()

	// unnamed root template
	var root = template.New("")

	var walkFunc = func(path string, info os.FileInfo, err error) (_ error) {

		// handle walking error if any
		if err != nil {
			return err
		}

		// skip all except regular files
		// TODO (kostyarin): follow symlinks
		if !info.Mode().IsRegular() {
			return
		}

		// filter by extension
		if filepath.Ext(path) != t.ext {
			return
		}

		// get relative path
		var rel string
		if rel, err = filepath.Rel(t.dir, path); err != nil {
			return err
		}

		// name of a template is its relative path
		// without extension
		rel = strings.TrimSuffix(rel, t.ext)

		// load or reload
		var (
			nt = root.New(rel)
			b  []byte
		)

		if b, err = ioutil.ReadFile(path); err != nil {
			return err
		}

		// necessary for reloading, this needs to come before parsing
		if t.funcs != nil {
			root = root.Funcs(t.funcs)
		}

		_, err = nt.Parse(string(b))
		return err
	}

	if err = filepath.Walk(t.dir, walkFunc); err != nil {
		return
	}

	t.Template = root // set or replace
	return
}

// IsModified lookups directory for changes to
// reload (or not to reload) templates if development
// pin is true.
func (t *Tmpl) IsModified() (yep bool, err error) {

	var errStop = errors.New("stop")

	var walkFunc = func(path string, info os.FileInfo, err error) (_ error) {

		// handle walking error if any
		if err != nil {
			return err
		}

		// skip all except regular files
		// TODO (kostyarin): follow symlinks
		if !info.Mode().IsRegular() {
			return
		}

		// filter by extension
		if filepath.Ext(path) != t.ext {
			return
		}

		if yep = info.ModTime().After(t.loadedAt); yep == true {
			return errStop
		}

		return
	}

	// clear the errStop
	if err = filepath.Walk(t.dir, walkFunc); err == errStop {
		err = nil
	}

	return
}

func (t *Tmpl) Render(w io.Writer, name string, data interface{}, c echo.Context) (err error) {
	log.Println("Ctx custom vrender initialized ...")

	// Add global methods if data is a map
	// if viewContext, isMap := data.(map[string]interface{}); isMap {
	// 	viewContext["reverse"] = c.Echo().Reverse
	// }

	// if devlopment
	if t.devel == true {

		// lookup directory for changes
		var modified bool
		if modified, err = t.IsModified(); err != nil {
			return
		}

		// reload
		if modified == true {
			if err = t.Load(); err != nil {
				return
			}
		}

	}

	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println("ctx error")
		log.Println(err)
	}
	return
}
