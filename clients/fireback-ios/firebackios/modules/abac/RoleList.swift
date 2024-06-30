//
//  RoleList.swift
//  firebackios
//
//  Created by ali on 12/01/2024.
//

import Foundation
import SwiftUI
import Pigeon
import Combine

struct RoleList: View {
    @ObservedObject var users = Query<Void, ArrayResponse<RoleEntity>>(
        key: QueryKey(value: "users"),
        behavior: .startImmediately(()),
        fetcher: {
            return GetRolesFetcher()
        }
    )
    
    var body: some View {
        var value = users.state.value
        
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
                            Text(item.name!)
                                .font(.system(size: 20, design: .default))
                            Text(item.name!)
                               .font(.system(size: 10, design: .default))
                        }
                        
                    }
                )
        }
    }
}

