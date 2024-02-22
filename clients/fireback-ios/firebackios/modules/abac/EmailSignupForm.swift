//
//  SignupForm.swift
//  firebackios
//
//  Created by ali on 10/01/2024.
//

import Foundation
import SwiftUI
import Combine
import Promises


struct PhoneSignupForm: View {
    @ObservedObject var vm = CheckClassicPassportActionReqDtoVm()
    @EnvironmentObject private var authService: AuthService


    enum NavDestination {
         case hasNoAccount
         case hasAccount
     }

    @State var destination : NavDestination?
    
    var body: some View {
        NavigationLink(destination: EmailSignupCompletionForm(initialEmail: vm.value), tag: NavDestination.hasNoAccount, selection: $destination) {
            EmptyView()
        }
        NavigationLink(destination: SigninClassicPasswordForm(initialEmail: vm.value), tag: NavDestination.hasAccount, selection: $destination) {
            EmptyView()
        }

        VStack(alignment: .leading) {
            H1(title: "Enter phone number")
            H2(content: "We'll check if you an account and provide next steps")
            
            VStack {
                FormPhoneNumber(label: "Phone Number", field: $vm.value, errorMessage: $vm.valueErrorMessage)
            }.padding(.bottom, 40).padding(.top, 40)
            
            
            AsyncButton(title: "Continue") {
                PostWorkspacePassportCheck(dto: vm.getDto())
            } onSuccess: { data in
                if data?.exists != nil && data!.exists! {
                    destination = .hasAccount
                } else {
                    destination = .hasNoAccount
                }
            } onFailure: { error in
                
            }
 
            Spacer()
        }
            .navigationBarTitle("Log in or sign up", displayMode: .inline)
            .padding(20)
    }
}

struct EmailSignupForm: View {
    @ObservedObject var vm = CheckClassicPassportActionReqDtoVm()
    @EnvironmentObject private var authService: AuthService


    enum NavDestination {
         case hasNoAccount
         case hasAccount
     }

    @State var destination : NavDestination?
    
    var body: some View {
        NavigationLink(destination: EmailSignupCompletionForm(initialEmail: vm.value), tag: NavDestination.hasNoAccount, selection: $destination) {
            EmptyView()
        }
        NavigationLink(destination: SigninClassicPasswordForm(initialEmail: vm.value), tag: NavDestination.hasAccount, selection: $destination) {
            EmptyView()
        }

        VStack(alignment: .leading) {
            H1(title: "Enter an email to get started")
            H2(content: "We'll check if you an account and provide next steps")
            
            VStack {
                FormEmail(label: "Email address", field: $vm.value, errorMessage: $vm.valueErrorMessage)
            }.padding(.bottom, 40).padding(.top, 40)
            
            
            AsyncButton(title: "Continue") {
                PostWorkspacePassportCheck(dto: vm.getDto())
            } onSuccess: { data in
                if data?.exists != nil && data!.exists! {
                    destination = .hasAccount
                } else {
                    destination = .hasNoAccount
                }
            } onFailure: { error in
                
            }
 
            Spacer()
        }
            .navigationBarTitle("Log in or sign up", displayMode: .inline)
            .padding(20)
    }
}

struct EnterOtpCodeScreen: View {
    @ObservedObject var vm = ClassicPassportOtpActionReqDtoVm()
    @EnvironmentObject private var authService: AuthService
    var initialValue: String?
    
    var body: some View {
        VStack(alignment: .leading) {
            H1(title: "Enter Otp Code")
            H2(content: "We just sent an activation code to you, type it here:")
            Spacer()
            H1(title: initialValue!)
            FormPassword(label: "Password", field: $vm.otp, errorMessage: $vm.otpErrorMessage)
            
            AsyncButton(title: "Continue") {
                PostWorkspacePassportOtp(dto: vm.getDto())
            } onSuccess: { data in
                if data?.session != nil {
                    authService.setSession(dto: data?.session)
                }
            } onFailure: { error in
                
            }
        }
        
            .navigationBarTitle("Otp", displayMode: .inline)
            .padding(20)
            .onAppear {
                vm.value = initialValue
            }
    }
}


struct SigninClassicPasswordForm: View {
    @ObservedObject var vm = ClassicSigninActionReqDtoVm()
    @EnvironmentObject private var authService: AuthService
    @State var secondScreen: Bool = false
    var initialEmail: String?

    
    enum NavDestination {
        case otp
    }
    
    @State var destination: NavDestination?
    
    var body: some View {
        NavigationLink(destination: EnterOtpCodeScreen(initialValue: initialEmail), tag: NavDestination.otp, selection: $destination) {
            EmptyView()
        }


        VStack(alignment: .leading) {
            H1(title: "Welcome back!")
            H2(content: "Enter your password to login")
            
            VStack {
                FormPassword(label: "Password", field: $vm.password, errorMessage: $vm.passwordErrorMessage)
            }.padding(.bottom, 40).padding(.top, 40)

            
            AsyncButton(title: "Continue") {
                PostPassportsSigninClassic(dto: vm.getDto())
            } onSuccess: { data in
                if data?.token != nil {
                    authService.setSession(dto: data)
                }
            } onFailure: { error in
                
            }
            
            
            AsyncTextButton(title: "Forgot password (OTP)") {
                return PostWorkspacePassportOtp(dto: ClassicPassportOtpActionReqDto(value: initialEmail))
            } onSuccess: { data in
                destination = .otp
            } onFailure: { error in
                destination = .otp
            }
            
             
             
            
            Spacer()
        }
            .navigationBarTitle("Log in", displayMode: .inline)
            .padding(20)
            .onAppear {
                vm.value = initialEmail
            }
    }
}


struct EmailSignupCompletionForm: View {
    @StateObject var vm = ClassicSignupActionReqDtoVm()
    @EnvironmentObject private var authService: AuthService
    var initialEmail: String?
    
    var body: some View {
        
        VStack(alignment: .leading) {
            
            
            VStack {
                FormText(label: "First name", field: $vm.firstName, errorMessage: $vm.firstNameErrorMessage)
                FormText(label: "Last name", field: $vm.lastName, errorMessage: $vm.lastNameErrorMessage)
                FormEmail(label: "Email", field: $vm.value, errorMessage: $vm.valueErrorMessage)
                FormPassword(label: "Password", field: $vm.password, errorMessage: $vm.passwordErrorMessage)
            }.padding(.bottom, 40).padding(.top, 40)
               
            Text("By selecting agree and continue below, I agree to Terms of Service, Payments Terms of Service, Privacy Policy, and Nondiscrimination Policy")
                
            
            AsyncButton<UserSessionDto>(title: "Continue") {
                var dto = vm.getDto()
                dto.type = .email
                return PostPassportsSignupClassic(dto: dto)
            } onSuccess: { data in
                authService.setSession(dto: data)
            } onFailure: { error in
                if (error?.errors) != nil {
                    
                    let err = error?.errors?.first(where: {$0.location == "firstName"})
                    vm.firstNameErrorMessage = err?.messageTranslated
                    
                    let err1 = error?.errors?.first(where: {$0.location == "lastName"})
                    vm.lastNameErrorMessage = err1?.messageTranslated
                    
                    let err2 = error?.errors?.first(where: {$0.location == "value"})
                    vm.valueErrorMessage = err2?.messageTranslated
                    
                    let err3 = error?.errors?.first(where: {$0.location == "password"})
                    vm.passwordErrorMessage = err3?.messageTranslated
                }
            }
            
            
            Spacer()
        }
            .navigationBarTitle("Finishing up signup", displayMode: .inline)
            .padding(20)
            .onAppear {
                vm.value = initialEmail
            }
    }
}
