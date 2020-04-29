package templates

import (
	"fmt"
	"strings"
	"text/template"
)

var ChartTemplate = `apiVersion: v1
description: A Helm chart for Kubernetes {{ .ReleaseName }}
name: {{ .ReleaseName }}
version: 1.0
`

var ServiceAccountTemplate = `apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .ReleaseName }}-{{ .Name }}
  labels:
    app.kubernetes.io/name: {{ .Name }}
	app.kubernetes.io/instance: {{ .ReleaseName }}
	app.kubernetes.io/version: {{ .Tag }}
    {{ if .Labels }}
	{{ range $key, $value := .Labels }}{{ $key }}: {{ $value }}
    {{ end }}
	{{ end }}
`

var ServiceTemplate = `apiVersion: v1
kind: Service
metadata:
	name: {{ .ReleaseName }}-{{ .Name }}
	labels:
		app.kubernetes.io/name: {{ .Name }}
		app.kubernetes.io/instance: {{ .ReleaseName }}
		app.kubernetes.io/version: {{ .Tag }}
		{{ if .Labels }}
		{{ range $key, $value := .Labels }}{{ $key }}: {{ $value }}
		{{ end }}
		{{ end }}
spec:
	type: ClusterIP
	ports:
	- name: http
	  port: 80
      targetPort: http
      protocol: TCP
	selector:
		app.kubernetes.io/name: {{ .Name }}
    	app.kubernetes.io/instance: {{ .ReleaseName }}
`

var DeploymentTemplate = `apiVersion: v1
kind: Deployment
metadata:
	name: {{ .ReleaseName }}-{{ .Name }}
	labels:
		app.kubernetes.io/name: {{ .Name }}
		app.kubernetes.io/instance: {{ .ReleaseName }}
		app.kubernetes.io/version: {{ .Tag }}
		{{ if .Labels }}
		{{ range $key, $value := .Labels }}{{ $key }}: {{ $value }}
		{{ end }}
		{{ end }}
spec:
	replicas: {{ .Replicas }}
	selector:
		matchLabels:
			app.kubernetes.io/name: {{ .Name }}
    		app.kubernetes.io/instance: {{ .ReleaseName }}
    template:
		metadata:
			labels:
				app.kubernetes.io/name: {{ .Name }}
    			app.kubernetes.io/instance: {{ .ReleaseName }}
			annotations:
		spec:
			serviceAccountName: {{ .ReleaseName }}-{{ .Name }}
			securityContext: 
			containers:
				- name: {{ .Name }}
                  image: {{ .Name}}:{{ .Tag}}
                  imagePullPolicy: IfNotPresent
                  {{ if .Command }}
                  command:{{ range $cmd := .Command }}
				   -  "{{$cmd}}"{{ end }} 
                  {{ end }}
                ports:
				  - name: http
                    containerPort: 8080
                    protocol: TCP

                livenessProbe:
                   httpGet:
                     path: {{ .LivenessProbe }}
                     port: http
                   initialDelaySeconds: 100
                   timeoutSeconds: 100                
                readinessProbe:
                   httpGet:
                     path: {{ .ReadinessProbe }}
                     port: http
                   initialDelaySeconds: 100
                   timeoutSeconds: 100

                resources:
                  limits:
                    cpu: "{{ index .Limits "cpu" }}"
                    memory:  "{{ index .Limits "memory" }}"   
                  requests:
                    cpu:  "{{ index .Limits "cpu" }}"
                    memory:  "{{ index .Limits "memory" }}"

                env:{{ range $key, $value := .EnvVars }}
                 - name: "{{ $key | ToUpper }}"
                   value: "{{ $value }}"{{end}}
			affinity:
			  nodeSelector:
			  tolerations:
`

//LoadTemplates parse static template to helm chart
func LoadTemplates(tName string, app *Application) *template.Template {
	switch tName {
	case "ChartTemplate":
		return getTemplate("Chart.yaml", ChartTemplate)
	case "DeploymentTemplate":
		return getTemplate(fmt.Sprintf("%s-deployment.yaml", app.Name), DeploymentTemplate)
	case "ServiceTemplate":
		return getTemplate(fmt.Sprintf("%s-service.yaml", app.Name), ServiceTemplate)
	case "ServiceAccountTemplate":
		return getTemplate(fmt.Sprintf("%s-serviceaccount.yaml", app.Name), ServiceAccountTemplate)
	}
	return nil
}

func getTemplate(name string, serviceTemplate string) *template.Template {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	tmpl, err := template.New(name).Funcs(funcMap).Parse(serviceTemplate)
	if err != nil {
		fmt.Println("Error parsing ", err)
	}
	return tmpl
}
