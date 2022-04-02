# SIAKADU INFROMATION EXTRACTOR

This app will extract information from SIAKADU which publicly available. This app will also provide utility to extract private information, but will require you to provide username and password of your SIAKADU account.

## Features

### 1. Get profile picture from NPM
- Command: **image**
- Args: NPM
- Result: Images from user with given NPM (if exists)
- Result location: *./result/images/NPM.jpg*
- Note: if NPM is not found from SIAKADU server, no result will be created.

### 2. Get Profile picture from NPM range
- Command: **image**
- Args: from (NPM) & to (NPM)
- Result: Images from user with NPM in given NPM range (start and to included)
- Result location *./result/images/start-to/NPM*
- Note: if NPM in given NPM range is not found from SIAKADU server, no result will be created.

### 3. Get public information form NPM
- Command: **public-info**
- Args: NPM
- Result: profile picture image, and a text containing information
- Result location: *./result/profile/NPM/*
- Note: if NPM is not found from SIAKADU server, no result will be created.

### 4. Get public information form NPM range
- Command: **public-info**
- Args: from (NPM) & to (NPM)
- Result: profile picture image, and a text containing information
- Result location: *./result/profile/from-to/NPM*
- Note: if NPM is not found from SIAKADU server, no result will be created.

### 5. Get All profile image
- Command: **images all**
- Result: all publicly available profile image
- Result location: *./images/full-${date}/*
- Note: if the NPM is not found, no image will be resulted.