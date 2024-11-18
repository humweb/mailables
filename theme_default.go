package mailables

type ThemeDefault struct {
}

func (t *ThemeDefault) Name() string {
	return "default"
}

func (t *ThemeDefault) Text() string {
	return ""
}
func (t *ThemeDefault) Html() string {
	return `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta name="color-scheme" content="light">
  <meta name="supported-color-schemes" content="light">
  <style rel="stylesheet" media="all">
    /* Base */

    body,
    body *:not(html):not(style):not(br):not(tr):not(code) {
      box-sizing: border-box;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif,
      'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol';
      position: relative;
    }

    body {
      -webkit-text-size-adjust: none;
      background-color: #ffffff;
      color: #718096;
      height: 100%;
      line-height: 1.4;
      margin: 0;
      padding: 0;
      width: 100% !important;
    }

    p,
    ul,
    ol,
    blockquote {
      line-height: 1.4;
      text-align: left;
    }

    a {
      color: #3869d4;
    }

    a img {
      border: none;
    }

    /* Typography */

    h1 {
      color: #3d4852;
      font-size: 18px;
      font-weight: bold;
      margin-top: 0;
      text-align: left;
    }

    h2 {
      font-size: 16px;
      font-weight: bold;
      margin-top: 0;
      text-align: left;
    }

    h3 {
      font-size: 14px;
      font-weight: bold;
      margin-top: 0;
      text-align: left;
    }
    .text-2xl {
      font-size: 24px;
      line-height: 32px;
    }
    .text-3xl{
      font-size: 30px;
      line-height: 36px;
    }
    .text-4xl{
      font-size: 36px;
      line-height: 40px;
    }
    p {
      font-size: 16px;
      line-height: 1.5em;
      margin-top: 0;
      text-align: left;
    }

    p.sub {
      font-size: 12px;
    }

    img {
      max-width: 100%;
    }

    /* Layout */

    .wrapper {
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      -premailer-width: 100%;
      background-color: #edf2f7;
      margin: 0;
      padding: 0;
      width: 100%;
    }

    .content {
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      -premailer-width: 100%;
      margin: 0;
      padding: 0;
      width: 100%;
    }

    /* Header */

    .header {
      padding: 25px 0;
      text-align: center;
    }

    .header a {
      color: #3d4852;
      font-size: 19px;
      font-weight: bold;
      text-decoration: none;
    }

    /* Logo */

    .logo {
      height: 35px;
      width: auto;
    }

    /* Body */

    .body {
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      -premailer-width: 100%;
      background-color: #edf2f7;
      border-bottom: 1px solid #edf2f7;
      border-top: 1px solid #edf2f7;
      margin: 0;
      padding: 0;
      width: 100%;
    }

    .inner-body {
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      -premailer-width: 95%;
      background-color: #ffffff;
      border-color: #e8e5ef;
      border-radius: 2px;
      border-width: 1px;
      box-shadow: 0 2px 0 rgba(0, 0, 150, 0.025), 2px 4px 0 rgba(0, 0, 150, 0.015);
      margin: 0 auto;
      padding: 0;
      width: 95%;
      border-radius: 10px;
    }

    /* Subcopy */

    .subcopy {
      border-top: 1px solid #e8e5ef;
      margin-top: 25px;
      padding-top: 25px;
    }

    .subcopy p {
      font-size: 14px;
    }

    /* Footer */

    .footer {
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      -premailer-width: 570px;
      margin: 0 auto;
      padding: 0;
      text-align: center;
      width: 570px;
    }

    .footer p {
      color: #b0adc5;
      font-size: 12px;
      text-align: center;
    }

    .footer a {
      color: #b0adc5;
      text-decoration: underline;
    }

    /* Tables */

    .table {
    border-collapse: collapse;
    margin: 25px 0;
    font-size: 0.9em;
    font-family: sans-serif;
    min-width: 400px;
	}
	.table thead tr {
		background-color: #6b7787;
		color: #ffffff;
		text-align: left;
	}
	.table th,
	.table td {
		padding: 12px 15px;
	}
	.table tbody tr {

	}
	
	.table tbody tr:nth-of-type(even) {
		background-color: #edf2f7;
	}
	
	.table tbody tr:last-of-type {
		border-bottom: 2px solid #94a3b8;
	}
	.table tbody tr.active-row {
		font-weight: bold;
		color: #94a3b8;
	}

    .content-cell {
      max-width: 100vw;
      padding: 32px;
    }

    /* Buttons */

    .action {
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      -premailer-width: 100%;
      margin: 30px auto;
      padding: 0;
      text-align: center;
      width: 100%;
      float: unset;
    }

    .btn {
      -webkit-text-size-adjust: none;
      border-radius: 4px;
      color: #fff;
      display: inline-block;
      overflow: hidden;
      text-decoration: none;
    }

    .btn-blue,
    .btn-primary {
      background-color: #2d3748;
      border-bottom: 8px solid #2d3748;
      border-left: 18px solid #2d3748;
      border-right: 18px solid #2d3748;
      border-top: 8px solid #2d3748;
    }

    .btn-green,
    .btn-success {
      background-color: #48bb78;
      border-bottom: 8px solid #48bb78;
      border-left: 18px solid #48bb78;
      border-right: 18px solid #48bb78;
      border-top: 8px solid #48bb78;
    }

    .btn-red,
    .btn-error {
      background-color: #e53e3e;
      border-bottom: 8px solid #e53e3e;
      border-left: 18px solid #e53e3e;
      border-right: 18px solid #e53e3e;
      border-top: 8px solid #e53e3e;
    }

    /* Panels */

    .panel {
      border-bottom: #d1d7dc solid 4px;
      margin: 21px 0;
      border-radius: 7px;
    }

    .panel-content {
      background-color: #edf2f7;
      color: #718096;
      padding: 16px;
      text-align: center;
      border-radius: 7px;
    }

    .panel-content p {
      color: #718096;
    }

    .panel-item {
      padding: 0;
    }

    .panel-item p:last-of-type {
      margin-bottom: 0;
      padding-bottom: 0;
    }

    /* Utilities */

    .break-all {
      word-break: break-all;
    }

    @media only screen and (max-width: 600px) {
      .inner-body {
        width: 100% !important;
      }

      .footer {
        width: 100% !important;
      }
    }

    @media only screen and (max-width: 500px) {
      .btn {
        width: 100% !important;
      }
    }
  </style>
</head>
<body>
{{- /*gotype: github.com/humweb/mailables.Page*/ -}}

<table class="wrapper" width="100%" cellpadding="0" cellspacing="0" role="presentation">
  <tr>
    <td align="center">
      <table class="content" width="100%" cellpadding="0" cellspacing="0" role="presentation">
        <tr>
          <td class="header">
            <a href="{{.Application.Link}}" target="_blank" style="display: inline-block;">
              {{ if .Application.Logo }}
              <img src="{{.Application.Logo | url }}" class="logo" alt="{{ .Application.Name }}" />
              {{ else }}
              {{ .Application.Name }}
              {{ end }}
            </a>
          </td>
        </tr>

        <!-- Email Body -->
        <tr>
          <td class="body" width="100%" cellpadding="0" cellspacing="0" style="border: hidden !important;">
            <table class="inner-body" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation">
              <!-- Body content -->

              <tr>
                <td class="content-cell">
                  {{ if .Body.Title }}
                  <h1>{{ .Body.Title }}</h1>
                  {{ else }}
                  {{ .Body.Greeting }} {{ .Body.Name }},<br><br>
                  {{ end }}

                  {{ with .Body.Intro }}
                  {{ .ToHTML }}
                  {{ end }}

                  {{ with .Body.Actions }}
                  {{ if gt (len . ) 0 }}
                  {{ range $action := . }}
                  {{/*<p>{{ $action.Instructions }}</p>*/}}
                  <table class="action" width="100%" cellpadding="0" cellspacing="0" role="presentation">
                    <tr>
            
					  <td width="100%" style="text-align: center;">
						<a href="{{ $action.Button.Link | url }}" class="btn btn-{{ $action.Button.Variant }}"
						   target="_blank" rel="noopener">{{ $action.Button.Text }}</a>
					  </td>
					</tr>
                  </table>
                  {{end}}
                  {{end}}
                  {{end}}
                  {{/*End Body Actions*/}}

                  <!-- Panels -->
                  {{with .Body.Panels}}
                  <table width="100%" cellpadding="0" cellspacing="0" role="presentation">
                    <tr>
                      {{ range $panel := . }}
                      <td style="padding: 6px;" width="{{$panel.Width}}">
                        <table class="panel" width="100%" cellpadding="0" cellspacing="0" role="presentation">
                          <tr>
                            <td class="panel-content">
                              <table width="100%" cellpadding="0" cellspacing="0" role="presentation">
                                <tr>
                                  <td class="panel-item">
                                    <h3>{{$panel.Title}}</h3>
                                    <p>{{$panel.Body.ToHTML}}</p>
                                  </td>
                                </tr>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                      {{end}}
                    </tr>
                  </table>
                  {{end}}

                  <!-- Table -->
                  {{ with .Body.Tables }}
                  {{ range $table := . }}
                  {{ $data := $table.Data }}
                  {{ if gt (len $data) 0 }}
                  <br>
                  <br>
                  <h2 class="text-2xl">{{ $table.Title }}</h2>

                  <table class="data-wrapper" width="100%" cellpadding="0" cellspacing="0">
                    <tr>
                      <td colspan="2">
                        <table class="table" width="100%" cellpadding="0" cellspacing="0">
                          <thead>
                          <tr>
                            {{ $col := index $data 0 }}
                            {{ range $entry := $col }}
                            <th>
                              {{ $entry.Key }}
                            </th>
                            {{ end }}
                          </tr>
                          </thead>
                          <tbody>
                          {{ range $row := $data }}
                          <tr>
                            {{ range $cell := $row }}
                            <td>
                              {{ $cell.Value }}
                            </td>
                            {{ end }}
                          </tr>
                          {{ end }}
                          </tbody>
                        </table>
                      </td>
                    </tr>
                  </table>
                  {{ end }}
                  {{ end }}
                  {{ end }}
                  {{ with .Body.Outro }}
                  <p>{{ .ToHTML }}</p>
                  {{ end }}
                  <p>
                    {{ .Body.Signature }},
                    <br />
                    {{ .Application.Name }}
                  </p>
                </td>
              </tr>
            </table>
          </td>
        </tr>
        <tr>
          <td>
            <table class="footer" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation">
              <tr>
                <td class="content-cell" align="center">
                  <p class="sub center">
                    {{ .Application.Copyright }}
                  </p>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>
</body>
</html>
`
}
