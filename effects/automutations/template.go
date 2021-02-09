package automutations

const templateString = `# --- {{.Name}}

input {{.Name}}Creator {
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

input {{.Name}}Update {
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
    update{{$name | title}}: [{{.Type}}Update!]
    {{- else}}
    {{$name}}: {{.Type}}SelectOrCreate
    {{- end}}
    {{- end}}
    {{- end}}
}

input {{.Name}}SelectOrCreate {
    fromID: ID
    create: {{.Name}}Creator
}

extend type Mutation {
    create{{.Name}}(data: {{.Name}}Creator!): {{.Name}}!
    createMany{{.PluralName}}(data: [{{.Name}}Creator!]!): [{{.Name}}!]!

    update{{.Name}}(data: {{.Name}}Update!): {{.Name}}!
    updateMany{{.PluralName}}(data: [{{.Name}}Update!]): [{{.Name}}!]!

    delete{{.Name}}(id: ID!): {{.Name}}!
    deleteMany{{.PluralName}}(ids: [ID!]!): [{{.Name}}!]!
}
`
