//
//  Forms.swift
//  firebackios
//
//  Created by ali on 17/01/2024.
//

import Foundation
import SwiftUI
import Combine

struct FormBaseWrapper: View {
    var field: Binding<String?>
    var body: some View {
        ZStack {
            RoundedRectangle(cornerRadius: 10)
                .strokeBorder(Color.gray, lineWidth: 1)
                .frame(height: 40)
            
            HStack {
                
                // Implement a way to show the children
                
                Button(action: {
                    field.wrappedValue = ""
                }) {
                    Image(systemName: "xmark.circle.fill")
                        .foregroundColor(.gray)
                        .opacity(field.wrappedValue == nil || field.wrappedValue!.isEmpty ? 0 : 1)
                }.padding(.trailing, 20)
            }
        }
    }
}

struct FormEmail: View {
    var label: String
    var field: Binding<String?>
    @Binding var errorMessage: String?
    
    var body: some View {
        ZStack {
            RoundedRectangle(cornerRadius: 10)
                .strokeBorder(Color.gray, lineWidth: 1)
                .frame(height: 40)
            
            HStack {
                TextField(label, text: field.toUnwrapped(defaultValue: ""))
                    .padding(.horizontal)
                    .keyboardType(.emailAddress)
                    .autocapitalization(.none)
                    .textInputAutocapitalization(.never)
//                    .disabled(disabled)
                Button(action: {
                    field.wrappedValue = ""
                }) {
                    Image(systemName: "xmark.circle.fill")
                        .foregroundColor(.gray)
                        .opacity(field.wrappedValue == nil || field.wrappedValue!.isEmpty ? 0 : 1)
                }.padding(.trailing, 20)
            }
            
            if let errorMessage = errorMessage {
                Text(errorMessage)
                    .font(.system(size: 13, design: .default))
                    .foregroundColor(.red)
            }
        }
    }
}


struct FormPhoneNumber: View {
    var label: String
    var field: Binding<String?>
    @Binding var errorMessage: String?
    
    var body: some View {
        ZStack {
            RoundedRectangle(cornerRadius: 10)
                .strokeBorder(Color.gray, lineWidth: 1)
                .frame(height: 40)
            
            HStack {
                TextField(label, text: field.toUnwrapped(defaultValue: ""))
                    .padding(.horizontal)
                    .keyboardType(.phonePad)
                    .autocapitalization(.none)
                    .textInputAutocapitalization(.never)
                Button(action: {
                    field.wrappedValue = ""
                }) {
                    Image(systemName: "xmark.circle.fill")
                        .foregroundColor(.gray)
                        .opacity(field.wrappedValue == nil || field.wrappedValue!.isEmpty ? 0 : 1)
                }.padding(.trailing, 20)
            }
            
            if let errorMessage = errorMessage {
                Text(errorMessage)
                    .font(.system(size: 13, design: .default))
                    .foregroundColor(.red)
            }
        }
    }
}

struct FormText: View {
   var label: String
   var field: Binding<String?>
   @Binding var errorMessage: String?
   var body: some View {
       ZStack {
           RoundedRectangle(cornerRadius: 10)
               .strokeBorder(Color.gray, lineWidth: 1)
               .frame(height: 50)
           
           HStack {
               TextField(label, text: field.toUnwrapped(defaultValue: ""))
                   .padding(.horizontal)
               Button(action: {
                   field.wrappedValue = ""
               }) {
                   Image(systemName: "xmark.circle.fill")
                       .foregroundColor(.gray)
                       .opacity(field.wrappedValue == nil || field.wrappedValue!.isEmpty ? 0 : 1)
               }.padding(.trailing, 20)
           }
           
           if let errorMessage = errorMessage {
               Text(errorMessage)
                   .font(.system(size: 13, design: .default))
                   .foregroundColor(.red)
           }
       }
   }
}


struct FormPassword: View {
   var label: String
   var field: Binding<String?>
   @Binding var errorMessage: String?
   var body: some View {
       ZStack {
           RoundedRectangle(cornerRadius: 10)
               .strokeBorder(Color.gray, lineWidth: 1)
               .frame(height: 50)
           
           HStack {
               SecureField(label, text: field.toUnwrapped(defaultValue: ""))
                   
                   .padding(.horizontal)
               Button(action: {
                   field.wrappedValue = ""
               }) {
                   Image(systemName: "xmark.circle.fill")
                       .foregroundColor(.gray)
                       .opacity(field.wrappedValue == nil || field.wrappedValue!.isEmpty ? 0 : 1)
               }.padding(.trailing, 20)
           }
           
           if let errorMessage = errorMessage {
               Text(errorMessage)
                   .font(.system(size: 13, design: .default))
                   .foregroundColor(.red)
           }
       }
   }
}

