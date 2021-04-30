# User Profile Demonstration
***
A simple application for viewing the user profiles and performing actions like create, update, read and delete of the user profiles. Powered with Vue.Js frontend and Golang as the backend server.

## Feature List
1. User can view the user profiles stored in the database.
2. User can create the user profiles to be stored in the database.
3. User can delete the user profiles stored in the database.
4. User can read the user profiles stored in the database.
5. User can sort the user profiles for effecient viewing.
6. User can search a specific user profile for effecient searching.

## Pre-requiste Steps To Follow
1. Please create a mongo database named as "Local_Database".
2. Please create a mongo collection named as "userProfile".
3. Please insert all the data from the file located at "user_profiles/static_content/Mock_Data_For_Database.json" into mongo database - "Local_Database", collection - "userProfile".
4. Please change the path variable "storagePath" inside the code for accessing your system local file path in the file - "/user_profiles/server/handlers/handler.go". Please check the location of the folder "user_profiles/server/App_File_Storage" and change according as it is the local static file server path.


## Execution of the project
1. Locate the client directory - "user_profile" inside the main project directory ("user_profiles/client/user_profile").
2. Execute the following command after installing the node modules using command - `npm install` or `yarn install` -

    ` npm run serve` or `yarn serve`
3. Locate the server directory inside the main project directory ("user_profiles/server").
4. Execute the following command -

    `go run main.go`
5. Both server and client will be running now. Please check on "http://localhost:8080" for the demonstration.

