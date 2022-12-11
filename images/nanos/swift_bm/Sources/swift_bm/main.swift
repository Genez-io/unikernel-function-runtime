import Foundation
let formatter = DateFormatter()
//2016-12-08 03:37:22 +0000
formatter.dateFormat = "yyyy-MM-dd HH:mm:ss Z"
let now = Date()
let dateString = formatter.string(from:now)
print(now)