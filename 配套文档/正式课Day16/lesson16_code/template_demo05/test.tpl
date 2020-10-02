Pet Info:
Name: {{ .Name }}
Age: {{ .Age }}
Owner: {{ .Owner.Name }}
{{ with .Owner }}
  Name: {{ .Name }}
  Age: {{ .Age }}
{{ end }}