package automutations

const templateString = `# --- {{.Name}}

{{- $creatorSuffix := "Creator" }}
{{- $updatorSuffix := "Updator" }}

{{- $creatorName := printf "%s%s" .Name $creatorSuffix }}
{{- $updatorName := printf "%s%s" .Name $updatorSuffix }}

input {{$creatorName}} {
    {{- range $name, $field := .ScalarFields}}
    {{- with $field}}
    {{$name}}: {{.Type}}{{if .Required}}!{{end}}
    {{- end}}
    {{- end}}
    {{ range $name, $field := .RelationFields}}
    {{- with $field}}
    {{- if .IsMultiple}}
    {{$name}}: [{{.Type}}SelectOrCreate!]{{if .Required}}!{{end}}
    {{- else}}
    {{$name}}: {{.Type}}SelectOrCreate{{if .Required}}!{{end}} 
    {{- end}}
    {{- end}}
    {{- end}}
}

input {{$updatorName}} {
    id: ID!

    {{- range $name, $field := .ScalarFields}}
    {{- with $field}}
    {{$name}}: {{.Type}}
    {{- end}}
    {{- end}}
    {{ range $name, $field := .RelationFields}}
    {{- with $field}}
    {{- if .IsMultiple}}
    add{{$name | title}}: [{{.Type}}SelectOrCreate!]
    delete{{$name | title}}: [ID!]
    update{{$name | title}}: [{{.Type}}{{$updatorSuffix}}!]
    {{- else}}
    {{$name}}: {{.Type}}SelectOrCreate
    {{- end}}
    {{- end}}
    {{- end}}
}

input {{.Name}}SelectOrCreate {
    fromID: ID
    create: {{$creatorName}}
}

extend type Mutation {
    create{{.Name}}(data: {{$creatorName}}!): {{.Name}}!
    createMany{{.PluralName}}(data: [{{$creatorName}}!]!): [{{.Name}}!]!

    update{{.Name}}(data: {{$updatorName}}!): {{.Name}}!
    updateMany{{.PluralName}}(data: [{{$updatorName}}!]): [{{.Name}}!]!

    delete{{.Name}}(id: ID!): {{.Name}}!
    deleteMany{{.PluralName}}(ids: [ID!]!): [{{.Name}}!]!
}
`
