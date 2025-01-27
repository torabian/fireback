//
//  action.swift
//  firebackios
//
//  Created by ali on 13/11/2023.
//

import Foundation


func doAction() {
    let url = URL(string: "https://www.stackoverflow.com")!

    let task = URLSession.shared.dataTask(with: url) {(data, response, error) in
        guard let data = data else { return }
        print(String(data: data, encoding: .utf8)!)
    }
    
    
    var m = UserEntity()

    
    
//    var m = AppMenuEntity()
//    m.LinkerId = "asdasd"
//    m.ParentId = "23123"
//    m.Href = "Href"
//    print(m)
//
//    let jsonEncoder = JSONEncoder()
//    let jsonData = try! jsonEncoder.encode(m)
//    let json = String(data: jsonData, encoding: String.Encoding.utf8)
//
//    
//    let drive = WidgetActions()
//
//    let decoder = JSONDecoder()
//
//    if let jsonPetitions = try? decoder.decode(AppMenuEntity.self, from: Data(json!.utf8)) {
//        print(jsonPetitions.ParentId)
//    }
//    task.resume()
}
