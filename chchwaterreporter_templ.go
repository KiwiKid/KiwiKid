// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func chchWaterReporter() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_2 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div id=\"chchwaterreporter\" class=\"min-h-screen\"><div class=\"w-full text-lg\"><div class=\"m-auto sm:w-4/5 max-w-6xl\"><a class=\"underline\" href=\"https://chch-water-reporter.vercel.app/\" target=\"_\">")
			if err != nil {
				return err
			}
			var_3 := `chch-water-reporter (live site)`
			_, err = templBuffer.WriteString(var_3)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
			var_4 := `provides a mobile friendly, representation of water usage in christchurch, its goal is to raise awareness for water conservation and the need for more broad understanding of how we consume our natural resources.`
			_, err = templBuffer.WriteString(var_4)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_5 := `With over 170,000 addresses on the map, the app required a dyanmic rendering system for higher zoom levels, which allowed for smooth scrolling, without compromissing on the detail display on the map.`
			_, err = templBuffer.WriteString(var_5)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><img src=\"./img/chch-reporting.PNG\" height=\"898\"> ")
			if err != nil {
				return err
			}
			var_6 := `[TODO: add more screen shots here]`
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Christchurch Water Reporter", "chchwaterreporter", false).Render(templ.WithChildren(ctx, var_2), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
