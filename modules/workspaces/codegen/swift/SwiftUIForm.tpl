import SwiftUI

struct {{ .e.FormName }}: View {

    @State private var formModel = {{ .e.EntityName }}()
    var body: some View {
        Form {
            Section {

                {{ range .e.CompleteFields }}

                {{if eq .Type "string"}}
                
                TextField("Email or Phone Number", text: $formModel.{{ .PublicName }}.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                {{ end }}
                {{ end }}
            }
        }
    }
}