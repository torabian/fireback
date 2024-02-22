//
//  ContentView.swift
//  firebackios
//
//  Created by ali on 13/11/2023.
//

import SwiftUI
import Eureka

struct User: Codable, Identifiable {
    let id: Int
    let name: String
}

struct UserFormModel {
    var username: String = ""
    var password: String = ""
    var isFormValid: Bool {
        return self.isValidUserName && self.isValidPassword
    }
    
    var isValidUserName : Bool {
        return !username.isEmpty
    }
    
    var isValidPassword: Bool {
        return password.count >= 6
    }
}

struct ContentView: View {
    @State private var formModel = UserFormModel()
    @EnvironmentObject private var authService: AuthService

    
    var body: some View {
        VStack {
            if authService.session != nil {
                TabView {
                    
                    NavigationView {
                        RoleList().navigationTitle("Roles")
                    }
                    .tabItem{
                        Image(systemName: "list.bullet")
                        Text("Roles")
                    }
                   
                    
                    NavigationView{
                        ProfileScreen().navigationTitle("Profile")
                    }
                    .tabItem {
                        Image(systemName: "person.crop.circle.fill")
                        Text("Profile")
                    }
                }
                
            } else {
                NavigationView {
                    ZStack {
//                        if authService.screen == .signin {
//                            NavigationLink(destination: SigninForm()) {
//                                Text("Go to detail")
////                                SigninForm().navigationBarBackButtonHidden(false)
//                            }
//                        }
//                        if authService.screen == .welcoome {
//
                            AuthenticationWelcomeScreen().navigationBarBackButtonHidden(false)
//                        }
                    }.animation(.spring(), value: authService.screen)
                }.navigationBarBackButtonHidden(false)
                
            }
        }

    }
}

 

struct ProfileScreen: View {
    @EnvironmentObject private var authService: AuthService
    var body: some View {
        Button(action: {
            authService.signOut()
        }) {
            Text("Logout")
        }
    }
}

struct RoundedButtonWithIcon: View {
    let imageName: String
    let buttonText: String
    
    var body: some View {
        Button(action: {
            // Navigate to signup screen
        }) {
            HStack {
                Image(systemName: imageName)
                    .foregroundColor(.white)
                Text(buttonText)
                    .foregroundColor(.white)
            }
            .padding()
            .background(Color.blue)
            .cornerRadius(20)
        }
    }
}


struct ScaleButtonStyle: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .scaleEffect(configuration.isPressed ? 0.9 : 1.0)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}

struct LoginScreen3: View {
    @State var name: String = ""
    @State var password: String = ""
    @State var showPassword: Bool = false

    var isSignInButtonDisabled: Bool {
     [name, password].contains(where: \.isEmpty)
    }

    var body: some View {
        VStack(alignment: .leading, spacing: 15) {
                    Spacer()
                    
                    TextField("Name",
                              text: $name ,
                              prompt: Text("Login").foregroundColor(.blue)
                    )
                    .padding(10)
                    .overlay {
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(.blue, lineWidth: 2)
                    }
                    .padding(.horizontal)

                RoundedButtonWithIcon(imageName: "google", buttonText: "Login with google")
                    HStack {
                        Group {
                            if showPassword {
                                TextField("Password", // how to create a secure text field
                                            text: $password,
                                            prompt: Text("Password").foregroundColor(.red)) // How to change the color of the TextField Placeholder
                            } else {
                                SecureField("Password", // how to create a secure text field
                                            text: $password,
                                            prompt: Text("Password").foregroundColor(.red)) // How to change the color of the TextField Placeholder
                            }
                        }
                        .padding(10)
                        .overlay {
                            RoundedRectangle(cornerRadius: 10)
                                .stroke(.red, lineWidth: 2) // How to add rounded corner to a TextField and change it colour
                        }

                        Button {
                            showPassword.toggle()
                        } label: {
                            Image(systemName: showPassword ? "eye.slash" : "eye")
                                .foregroundColor(.red) // how to change image based in a State variable
                        }

                    }.padding(.horizontal)

                    Spacer()

                    Button {
                        print("do login action")
                    } label: {
                        Text("Sign In")
                            .font(.title2)
                            .bold()
                            .foregroundColor(.white)
                    }
                    .frame(height: 50)
                    .frame(maxWidth: .infinity) // how to make a button fill all the space available horizontaly
                    .background(
                        isSignInButtonDisabled ? // how to add a gradient to a button in SwiftUI if the button is disabled
                        LinearGradient(colors: [.gray], startPoint: .topLeading, endPoint: .bottomTrailing) :
                            LinearGradient(colors: [.blue, .red], startPoint: .topLeading, endPoint: .bottomTrailing)
                    )
                    .cornerRadius(20)
                    .disabled(isSignInButtonDisabled) // how to disable while some condition is applied
                    .padding()
                }
    }
}

struct FirstScreen: View {
    @State private var formModel = UserFormModel()
    var body: some View {
        Form {
            Section {
                TextField("Username", text: $formModel.username)
                    .autocapitalization(.none)
                if !formModel.isValidUserName {
                    Text("Please enter a valid username")
                        .foregroundColor(.red)
                }
                
                
                SecureField("Password", text: $formModel.password)
                if !formModel.isValidPassword {
                    Text("Please enter a valid password")
                        .foregroundColor(.red)
                }
                
            }
        }
    }
}

struct SecondScreen: View {
    var body: some View {
        Text("This is the second screen")
    }
}
