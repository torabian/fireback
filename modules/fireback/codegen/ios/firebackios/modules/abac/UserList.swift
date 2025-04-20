//
//  UserList.swift
//  firebackios
//
//  Created by ali on 12/01/2024.
//

import Foundation
import SwiftUI
import Pigeon
import Combine

struct UserList: View {
    @ObservedObject var users = Query<Void, ArrayResponse<UserEntity>>(
        key: QueryKey(value: "users"),
        behavior: .startImmediately(()),
        fetcher: {
            return GetUsersFetcher()
        }
    )
    
    
    var body: some View {
        var value = users.state.value
        
        print(users.state.value)
        switch users.state {
            case .idle:
                return AnyView(Text("Ide..."))
            case .loading:
                return AnyView(Text("Loading..."))
            case .failed:
                return AnyView(Text("Oops..."))
            case .succeed:
                return AnyView(
                    List(value?.data?.items ?? []) { item in
                        VStack {
                            HStack {
                                Text(item.person?.firstName! ?? "-")
                                    .font(.system(size: 20, design: .default))
                                Text(item.person?.lastName! ?? "-")
                                    .font(.system(size: 10, design: .default))
                                
                            }
                        }
                        
                    }
                )
        }
    }
}

