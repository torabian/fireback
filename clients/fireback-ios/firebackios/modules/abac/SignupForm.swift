//
//  SignupForm.swift
//  firebackios
//
//  Created by ali on 10/01/2024.
//

import Foundation
import SwiftUI
import Combine


struct SignupForm: View {
 
    @State public var firstName = ""
    @State private var lastName = ""
    @State private var birthDate = Date()
    @State private var email = "admin"
    @State private var password = "admin"
    @State private var isSecureEntry = true
 
    var body: some View {
        
        Form {
            TextField("Email address", text: $email)
            TextField("First Name", text: $firstName)
            TextField("Last Name", text: $lastName)
            DatePicker("Birthday", selection: $birthDate, displayedComponents: .date)
            
            HStack {
                if isSecureEntry {
                    SecureField("Password", text: $password)
                } else {
                    TextField("Password", text: $password)
                }
                
                Button(action: {
                    isSecureEntry.toggle()
                }) {
                    Image(systemName: "eye")
                }
            }
            
            Button(action: {
                PostPassportSigninEmail(dto: EmailAccountSigninDto(email: email, password: password)).then { res in
                    if res != nil && res?.token != nil {
//                        authservice
                    }
                    
                }.catch{  reject in
                    print(reject)
                }
                
            }) {
                HStack {
                    Spacer()
                    Text("Submit")
                        .foregroundColor(.white)
                    Spacer()
                }
            }
            .padding()
            .background(Color.blue)
            .cornerRadius(10)
        }
    }
}
