# `PROJECT'S TITLE:`  🎧GROUPIE~TRACKER-SEARCH-BAR🎧

## `TABLE OF CONTENTS🎶`
- [🎯 OBJECTIVES](#🎯objectives)
- [📑DESCRIPTION](#📑description)
- [🛠️FEATURES](#🛠️features)
- [🧭HOW_TO_RUN](#🧭how_to_run)
- [MAIN-INTERFACE](#🖥️interface)
- [👤AUTHORS](#👤authors)

## 🎯`OBJECTIVES:`
Groupie tracker search bar consists of creating a functional program that searches, inside your website, for a specific text input. So the focus of this project is to create a way for the client to search a member or artist or any other attribute in the data system you made.

<!-- It will be given an API, that consists in four parts:
   - 📀Artists
   - 📍Locations
  -  🗓️Dates
  -  🎫Relation.
 -->
  ## 📑`DESCRIPTION:`
- [x] The program  handle these search cases :
      - [] artist/band name
      - []  members
      - [] locations
      - [] first album date
      - [] creation date
- [x] The program handle search input as case-insensitive.
- [x] The search bar  have typing suggestions as you write.
       - [] The search bar identify and display in each suggestion the individual type of the search cases. (ex: japan -> location)
      -  For example if you start writing "phil" it should appear as suggestions Phil Collins - member and Phil Collins - artist/band. This is just an example of a display.

### 🛠️`FEATURES:`

- Rich visualizations: explore the details through well-organized maps of your favorite artists
 <!-- - Personalization: Customize your experience by filtering results by year, music genre or location. -->
-  User-friendly: An intuitive, interactive user interface makes navigation easy, providing a seamless experience for music fans.
- Real-time updates: Enjoy instant updates on artists, venues and dates thanks to dynamic API integration.

#### search features

- [x] Opening the Search Bar :
       - [] Click on the search icon represented by the magnifying glass
  - [x]  Entering the Search :
       - [] Once the search bar is open, start typing your query in the text box provided. 
  - [x]  Viewing results :
      - []  Search results will be displayed in real time as you type.
  - [x]  Quick Navigation:
      - []  Click on a result to be taken instantly to the corresponding section of the page. The search bar will close automatically.

  - [x]  Closing the Search Bar :
      - [] To close the search bar, click outside it or use the close button.

## 🧭`HOW_TO_RUN`
### 🐾`TO RUN IT:`
_Make sure you're in the right folder, and set to terminal the command below to Start the web server:_
```bash
go run .
```
_Access the web application in your browser_
```bash
🌐 http://localhost:8080
```
### 📇`Summary of Steps` :

1️⃣ Users click on see artists to see a map of all the artists on the platform.

2️⃣ once in the artist interface.
   - 🔢 find your artist 
   - 🔢 When the user clicks on the button "voir":
       - artist information is displayed .
       - click on the cross icon or just outside to close the artist info panel
   - 🔢 The user can repeat the process with new card or a different artist.
   <!-- - 🔢 -->

### 🚷`Error Handling`
The web server handles various error scenarios (errors that have been taken into account and should be avoided ), such as:
- ⚠️ 404 Not Found: Page not found.
- ⚠️ 400 Bad Request: Invalid request.
- ⚠️ 405 Method Not Allowed: HTTP request with an unauthorized method
- ⚠️ 500 Internal Server Error: Server-side errors.

## 🖥️`INTERFACE`


![Main Interface](/vues/assets/img/homeinterface.png)

<hr>

![Main Interface](/vues/assets/img/homeinterface2.png)
<hr>
Interface de la page d'accueil GROUPIE TRACKER
<hr>

## 👤`AUTHORS`
+ 🧑[@mabdelwe](https://learn.zone01dakar.sn/git/mabdelwe)
+ 🧑[@lodiouf](https://learn.zone01dakar.sn/git/lodiouf)
+ 🧑[@jgoudiab](https://learn.zone01dakar.sn/git/jgoudiab)