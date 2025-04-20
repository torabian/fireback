//
//  Buttons.swift
//  firebackios
//
//  Created by ali on 17/01/2024.
//

import Foundation
import Promises
import SwiftUI

struct PrimaryButton: View {
    var title: String
    @Binding var disabled: Bool
    var action: () -> Void
    
    var body: some View {
        Button(action: action) {
            HStack {
                Spacer()
                Text(title)
                    .foregroundColor(.white)
                Spacer()
            }
        }
        .padding()
        .background(disabled ? Color.gray : Color.blue)
        .cornerRadius(25)
        .disabled(disabled)
    }
}

struct AsyncButton<T>: View {
    var title: String
    var action: () -> Promise<T?>
    var onSuccess: (_ data: T?) -> Void
    var onFailure: (_ error: IResponseError?) -> Void
    
    @State private var isLoading = false
    @State private var errorMessage: String?
    
    var body: some View {
        ZStack {
           
            PrimaryButton( title: title, disabled: $isLoading) {
                isLoading = true
                action().then { res in
                    isLoading = false
                    onSuccess(res)
                }.catch{reject in
                   isLoading = false
                   if let customError = reject as? IResponseError {
                       onFailure(customError)
                       errorMessage = customError.messageTranslated
                   } else {
                       errorMessage = "Unknown error has happened"
                       onFailure(nil)
                   }
                }
                
                
            }
            if isLoading {
                ProgressView()
                    .progressViewStyle(CircularProgressViewStyle())
                    .padding()
                    .offset(x: -150)
            }
            
        }
        if let errorMessage = errorMessage {
            Text(errorMessage)
                .font(.system(size: 13, design: .default))
                .foregroundColor(.red)
        }

    }
}


struct AsyncTextButton<T>: View {
    var title: String
    var action: () -> Promise<T?>
    var onSuccess: (_ data: T?) -> Void
    var onFailure: (_ error: IResponseError?) -> Void
    
    @State private var isLoading = false
    @State private var errorMessage: String?
    
    var body: some View {
        ZStack {
           
            Button() {
                isLoading = true
                action().then { res in
                    isLoading = false
                    onSuccess(res)
                }.catch{reject in
                   isLoading = false
                   if let customError = reject as? IResponseError {
                       onFailure(customError)
                       errorMessage = customError.messageTranslated
                   } else {
                       errorMessage = "Unknown error has happened"
                       onFailure(nil)
                   }
                }
                
                
            } label: {
                Text(title)
            }
            if isLoading {
                ProgressView()
                    .progressViewStyle(CircularProgressViewStyle())
                    .padding()
                    .offset(x: -150)
            }
            
        }
        if let errorMessage = errorMessage {
            Text(errorMessage)
                .font(.system(size: 13, design: .default))
                .foregroundColor(.red)
        }

    }
}
