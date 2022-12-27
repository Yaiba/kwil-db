{{/*
Expand the name of the chart.
*/}}
{{- define "hasura.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "hasura.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "hasura.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "hasura.labels" -}}
helm.sh/chart: {{ include "hasura.chart" . }}
{{ include "hasura.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/component: "graphql"
{{- end }}

{{/*
Selector labels
*/}}
{{- define "hasura.selectorLabels" -}}
app.kubernetes.io/name: {{ include "hasura.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "hasura.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "hasura.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}


{{/*
initContainers for database depencency
*/}}
{{- define "hasura.initContainers.db" -}}
- name: init-wait-db-service
  image: busybox:1.35.0
  command: 
    - sh
    - -c
    - |
      {{- $service := include "postgresql.primary.fullname" .Subcharts.postgresql }}
      echo try resolve dns name "{{ $service }}"
      until nslookup {{ $service }}.$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace).svc.cluster.local
      do 
        echo waiting for {{ $service }}
        sleep 2
      done        
- name: init-wait-postgres
  image: busybox:1.35.0
  envFrom:
  - configMapRef:
      name: {{ include "hasura.fullname" . }}
  command: 
    - sh
    - -c
    - |
      {{- $host := include "postgresql.primary.fullname" .Subcharts.postgresql }}
      {{- $port := .Values.postgresql.primary.service.ports.postgresql }}
      echo scan {{ $host }}:{{ $port }}
      for i in $(seq 1 300); 
      do 
        nc -zvw1 {{ $host}} {{ $port }} && exit 0 || sleep 3
      done
      exit 1
{{- end}}