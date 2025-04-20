//
//  SignupForm.swift
//  firebackios
//
//  Created by ali on 10/01/2024.
//

import Foundation
import SwiftUI
import Combine


struct ImageTouchable: View {
    
    var image: String
    var label: String
    var isSystemImage: Bool
    
    var body: some View {
        HStack {
            if isSystemImage {
                Image(systemName: image)
                    .resizable()
                    .aspectRatio(contentMode: .fit)
                    .frame(width: 20, height: 20)
            } else {
                Image(image)
                    .resizable()
                    .aspectRatio(contentMode: .fit)
                    .frame(width: 20, height: 20)
            }
           
            
            Spacer()
            Text(label)
                .font(.headline)
                .fontWeight(.bold)
            
            Spacer()
        }
        .padding()
        .foregroundColor(.black)
        .frame(maxWidth: .infinity)
        .overlay(
                  RoundedRectangle(cornerRadius: 30)
                      .stroke(Color.black, lineWidth: 1)
              )
    }
}


struct ImageButton: View {
    
    var image: String
    var label: String
    var isSystemImage: Bool
    var action: () -> Void
    
    var body: some View {
            Button(action: action) {
                ImageTouchable(image: image, label: label, isSystemImage: isSystemImage)
            }
        }
}


struct H1 : View {
    var title: String
    var body: some View {
        HStack {
            Text(title)
                .font(.system(size: 24, weight: .bold, design: .default))
                .padding(.bottom, 20)
        }.frame(
            maxWidth: .infinity,
            alignment: .leading
        )
        
    }
}


struct H2 : View {
    var content: String
    var body: some View {
        HStack {
            Text(content)
                .font(.system(size: 16,  design: .default))
                .padding(.bottom, 20)
        }.frame(
            maxWidth: .infinity,
            alignment: .leading
        )
        
    }
}



struct AuthenticationWelcomeScreen: View {
    
    @EnvironmentObject private var authService: AuthService
    
    var body: some View {
        VStack(alignment: .center) {
            Spacer().frame(maxHeight: .infinity)
            
            H1(title: "Signin and keep your data")
            H2(content: "Signing up using any of the methods below will help you to keep the information in the cloud, you can use them also if you have already an account.")
            
            Spacer().frame(maxHeight: .infinity)
            VStack (alignment: .leading, spacing: 15 ){
                ImageButton(image: "google", label: "Continue with Google", isSystemImage: false) {
                    print("Implement")
                }
                NavigationLink(destination: SigninForm()) {
                    ImageTouchable(image: "faceid", label: "Continue with Apple", isSystemImage: true)
                }
                ImageButton(image: "facebook", label: "Continue with Facebook",  isSystemImage: false) {
                    print("Implement")
                }

                NavigationLink(destination: PhoneSignupForm()) {
                    ImageTouchable(image: "phone.circle", label: "Continue with Phone" , isSystemImage: true)
                }

                
//                ImageButton(image: "phone.circle", label: "Continue with Phone" , isSystemImage: true) {
//                    print("Implement")
//                }
                
                // This is using navigation link
                NavigationLink(destination: EmailSignupForm()) {
                    ImageTouchable(image: "envelope", label: "Continue with Email" , isSystemImage: true)
                }
                 

                Spacer()
            }
            .frame(maxWidth: .infinity, // Full Screen Width
                   maxHeight: .infinity, // Full Screen Height
                   alignment: .topLeading)
            
            
        }.padding(25).frame(maxWidth: .infinity, // Full Screen Width
                            maxHeight: .infinity, // Full Screen Height
                            alignment: .topLeading)

    }
}
