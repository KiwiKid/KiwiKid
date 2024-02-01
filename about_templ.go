// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func about(yearSince string) templ.Component {
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
			_, err = templBuffer.WriteString("<div class=\"w-full text-center py-2\"><p class=\"italic text-gray-600\">")
			if err != nil {
				return err
			}
			var_3 := `(As a Kiwi, this level of self-promotion makes me deeply uncomfortable, but here we go...)`
			_, err = templBuffer.WriteString(var_3)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p></div> <p class=\"mb-4 text-gray-700\">")
			if err != nil {
				return err
			}
			var_4 := `I have been developing in a professional capacity as a full stack developer for over `
			_, err = templBuffer.WriteString(var_4)
			if err != nil {
				return err
			}
			var var_5 string = yearSince
			_, err = templBuffer.WriteString(templ.EscapeString(var_5))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_6 := `years.`
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p class=\"mb-4 text-gray-700\">")
			if err != nil {
				return err
			}
			var_7 := `I have extensive experience in implementing complex software solutions, based on varying degrees of specification.`
			_, err = templBuffer.WriteString(var_7)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p class=\"mb-4 text-gray-700\">")
			if err != nil {
				return err
			}
			var_8 := `Working closely with stakeholders, refining and architecting solutions through to deployment, maintenance and support.`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p class=\"mb-4 text-gray-700\">")
			if err != nil {
				return err
			}
			var_9 := `I am passionate about solving real problems, working with the best tools (or getting there) and improving the overall effectiveness of the organisation and the team`
			_, err = templBuffer.WriteString(var_9)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <p class=\"mb-4 text-gray-700\">")
			if err != nil {
				return err
			}
			var_10 := `I am excited about continuing to explore the benefits and opportunities of serverless/cloud solutions.`
			_, err = templBuffer.WriteString(var_10)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p> <div class=\"text-center w-full\"><a class=\"underline text-2xl hover:text-blue-500 font-semibold\" href=\"mailto:ghcumming01@gmail.com\">")
			if err != nil {
				return err
			}
			var_11 := `Get in touch`
			_, err = templBuffer.WriteString(var_11)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Me", "intro").Render(templ.WithChildren(ctx, var_2), templBuffer)
		if err != nil {
			return err
		}
		var_12 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"border border-black p-4\"><div class=\"grid grid-cols-3 space-y-2 sm:grid-cols-3 sm:space-x-4\"><div class=\"\"><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_13 := `Loving`
			_, err = templBuffer.WriteString(var_13)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div class=\"italic text-sm\">")
			if err != nil {
				return err
			}
			var_14 := `Development`
			_, err = templBuffer.WriteString(var_14)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_15 := `Go, a-h/templ, htmx, C#, .Net Core + DI, React, Tailwind, DynamoDB, Swagger, Storybook, postgres `
			_, err = templBuffer.WriteString(var_15)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div class=\"italic text-sm\">")
			if err != nil {
				return err
			}
			var_16 := `Software`
			_, err = templBuffer.WriteString(var_16)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_17 := `Home Assistant, Frigate, Node-Red, Heroku`
			_, err = templBuffer.WriteString(var_17)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div><div><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_18 := `Learning`
			_, err = templBuffer.WriteString(var_18)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_19 := `Go, Fiber`
			_, err = templBuffer.WriteString(var_19)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div><div><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_20 := `Up Next`
			_, err = templBuffer.WriteString(var_20)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_21 := `Svelte, Unraid, `
			_, err = templBuffer.WriteString(var_21)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div></div><div class=\"col-span-full\"><div class=\"text-xl text-center\">")
			if err != nil {
				return err
			}
			var_22 := `Experience In`
			_, err = templBuffer.WriteString(var_22)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div>")
			if err != nil {
				return err
			}
			var_23 := `Go, C#, postgres, Tailwind, postgres, Next.js + Vercel, SQL, SQL Server, Node.js, New Relic, T-SQL, ASP.NET MVC, Web APIs, IIS, AWS CDK, RESTful API design, git, CI/CD, AWS Lambda, VSCode, Heroku, Entity Framework`
			_, err = templBuffer.WriteString(var_23)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Tech", "tech").Render(templ.WithChildren(ctx, var_12), templBuffer)
		if err != nil {
			return err
		}
		var_24 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_25 := `Luxury Escapes is one of the leading travel provider in Australia.....`
			_, err = templBuffer.WriteString(var_25)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Luxury Escapes", "lux").Render(templ.WithChildren(ctx, var_24), templBuffer)
		if err != nil {
			return err
		}
		var_26 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_27 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				var_28 := `Streamliners manages several company platforms, including the largest, Healthpathways, which comprises of a responsive websites, publishing platform, and a shared administration and feedback system serving 30 million patients.`
				_, err = templBuffer.WriteString(var_28)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Background").Render(templ.WithChildren(ctx, var_27), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_29 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				var_30 := `React, C#, .Net Core, ASP.NET Core Web APIs, Entity Framework, TeamCity, GitHub Actions, SQL, Unit Testing, Dependency Injection, ELK, ASP.NET Web Forms, AWS CDK + Range of AWS services including DynamoDB, Lambda, S3 etc`
				_, err = templBuffer.WriteString(var_30)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Tech Used").Render(templ.WithChildren(ctx, var_29), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_31 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				var_32 := `Whilst at Streamliners, a major focus of the work has been a move towards a cloud-based microservice based architecture.`
				_, err = templBuffer.WriteString(var_32)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_33 := `I've been really engaged by exploring the benefits and opportunities of serverless solutions for quickly developing really complex software solutions. Leading the design and implementation of these technologies gave me a first hand look at the advantages (and disadvantages!) of building and deploying cloud solutions.`
				_, err = templBuffer.WriteString(var_33)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_34 := `I've gain a huge respect for the potential for these tools to empower development teams and, quite honestly, its excites me greatly.`
				_, err = templBuffer.WriteString(var_34)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Focus").Render(templ.WithChildren(ctx, var_31), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_35 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				var_36 := `Working with the team to improve effectiveness & apply continuous improvement principles`
				_, err = templBuffer.WriteString(var_36)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_37 := `Building Process Jira Workflows & Dashboards for the Team`
				_, err = templBuffer.WriteString(var_37)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_38 := `Running Retros`
				_, err = templBuffer.WriteString(var_38)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_39 := `Working with stakeholders to remove blockers`
				_, err = templBuffer.WriteString(var_39)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_40 := `Proposing and developing solutions for several company platforms`
				_, err = templBuffer.WriteString(var_40)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Scrum Master").Render(templ.WithChildren(ctx, var_35), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Streamliners", "pro").Render(templ.WithChildren(ctx, var_26), templBuffer)
		if err != nil {
			return err
		}
		var_41 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_42 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				var_43 := `eStar is one of Australasia’s leading e-commerce SaaS platforms creating high-end bespoke e-commerce websites for large international companies.`
				_, err = templBuffer.WriteString(var_43)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Background").Render(templ.WithChildren(ctx, var_42), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_44 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				var_45 := `IIS, C#, SQL, .Net, ASP.NET Core Web APIs, Unit Testing, ASP.NET Web Forms`
				_, err = templBuffer.WriteString(var_45)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Tech Used").Render(templ.WithChildren(ctx, var_44), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_46 := `While at eStar built various system components and learnt various technologies.`
			_, err = templBuffer.WriteString(var_46)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_47 := `Development spanned the entire stack, front-end web development for client websites as well as the iSAMS administration portal.`
			_, err = templBuffer.WriteString(var_47)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_48 := `Role`
			_, err = templBuffer.WriteString(var_48)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_49 := `IIS, C#, SQL, .Net, ASP.NET Core Web APIs, Unit Testing, ASP.NET Web Forms`
			_, err = templBuffer.WriteString(var_49)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_50 := `While at eStar built various system components and learnt various technologies.`
			_, err = templBuffer.WriteString(var_50)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_51 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<div>")
				if err != nil {
					return err
				}
				var_52 := `Development spanned the entire stack, front-end web development for client websites as well as the federated muli-client iSAMS administration portal.`
				_, err = templBuffer.WriteString(var_52)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_53 := `I had a number of opportunities at eStar, giving me a wide range of development experience across a number of teams:`
				_, err = templBuffer.WriteString(var_53)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_54 := `David Jones Project Team`
				_, err = templBuffer.WriteString(var_54)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_55 := `Working with large multinational clients`
				_, err = templBuffer.WriteString(var_55)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_56 := `Meeting complex stakeholder goals`
				_, err = templBuffer.WriteString(var_56)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_57 := `Developing and optimising for an application operating at a very large scale (~2.5 M Visits/mth)`
				_, err = templBuffer.WriteString(var_57)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_58 := `Air New Zealand`
				_, err = templBuffer.WriteString(var_58)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_59 := `Leading meetings with clienta`
				_, err = templBuffer.WriteString(var_59)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_60 := `Design, developed, and supported a highly custom multi-tenancy external supplier product management system requiring wide-scale refactoring of a large existing e-commerce solution`
				_, err = templBuffer.WriteString(var_60)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_61 := `On Call Support (24/7)`
				_, err = templBuffer.WriteString(var_61)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_62 := `Triage client on-call requests`
				_, err = templBuffer.WriteString(var_62)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_63 := `Addressing and escalate as required`
				_, err = templBuffer.WriteString(var_63)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				var_64 := `Managing the response and writing root cause incident reports`
				_, err = templBuffer.WriteString(var_64)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div>")
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = subbox("Tech Used").Render(templ.WithChildren(ctx, var_51), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("eStar", "pro").Render(templ.WithChildren(ctx, var_41), templBuffer)
		if err != nil {
			return err
		}
		var_65 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"text-center font-medium text-gray-800\">")
			if err != nil {
				return err
			}
			var_66 := `I love building and flying First-Person-View Quadcopters.`
			_, err = templBuffer.WriteString(var_66)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div> <div class=\"text-center font-medium text-gray-800\">")
			if err != nil {
				return err
			}
			var_67 := `(The photo of me at the top is after clambering up that mountain to retrieve a lost quad!)`
			_, err = templBuffer.WriteString(var_67)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div> <img src=\"./img/atopthehill.webp\"> <div class=\"italic text-center text-gray-600\">")
			if err != nil {
				return err
			}
			var_68 := `Watch out for YouTube compression here!`
			_, err = templBuffer.WriteString(var_68)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div> <div class=\"m-auto\"><div style=\"width: 100dhv\" class=\"min-h-4\"><div style=\"width: 100%; height: 100%;\"><iframe allowfullscreen=\"\" allow=\"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share\" title=\"Measured Manoeuvres\" src=\"https://www.youtube.com/embed/fxbgoZZdHD4?autoplay=0&amp;mute=0&amp;controls=0&amp;origin=https%3A%2F%2Fgregc.dev&amp;playsinline=1&amp;showinfo=0&amp;rel=0&amp;iv_load_policy=3&amp;modestbranding=true&amp;loop=1&amp;enablejsapi=1&amp;widgetid=1\" id=\"widget2\" width=\"100%\" height=\"100%\" frameborder=\"0\"></iframe></div></div></div> <div class=\"text-xl text-center font-medium text-gray-800\">")
			if err != nil {
				return err
			}
			var_69 := `something more scenic....`
			_, err = templBuffer.WriteString(var_69)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div> <div class=\"m-auto\"><div style=\"width: 100dhv\" class=\"min-h-4\"><div style=\"width: 100%; height: 100%;\"><iframe frameborder=\"0\" allowfullscreen=\"\" allow=\"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share\" title=\"Arthur&#39;s Pass\" width=\"100%\" height=\"100%\" src=\"https://www.youtube.com/embed/CaHXT2S6gJA?autoplay=0&amp;mute=0&amp;controls=0&amp;origin=https%3A%2F%2Fgregc.dev&amp;playsinline=1&amp;showinfo=0&amp;rel=0&amp;iv_load_policy=3&amp;modestbranding=true&amp;loop=1&amp;enablejsapi=1&amp;widgetid=3\" id=\"widget4\"></iframe></div></div></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Flying Drones FPV (First Person View)", "fpv").Render(templ.WithChildren(ctx, var_65), templBuffer)
		if err != nil {
			return err
		}
		var_70 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div>")
			if err != nil {
				return err
			}
			var_71 := `[INSERT FIXED WING WORDS + PICS]`
			_, err = templBuffer.WriteString(var_71)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("RC Model Airplanes", "fixed-fpv").Render(templ.WithChildren(ctx, var_70), templBuffer)
		if err != nil {
			return err
		}
		var_72 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div>")
			if err != nil {
				return err
			}
			var_73 := `[INSERT HOME-LAB]`
			_, err = templBuffer.WriteString(var_73)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = box("Smart home/Homelab", "home-lab").Render(templ.WithChildren(ctx, var_72), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
