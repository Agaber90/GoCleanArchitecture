# Screenshot using Golang with clean architecture
This a Golang project to start take a screenshot form URLS using pdfcrowd API and save the images data to mysql database
before start to explain this porject please

# install the following packages 
<br />
<ol>
  
  <li> go get github.com/go-sql-driver/mysql </li>
  <li> go get github.com/labstack/echo </li>
  <li> go get -u github.com/spf13/viper </li>
  <li> go get github.com/pdfcrowd/pdfcrowd-go</li>
  </ol>
<br />

<p>First of all let me speak about the solution architecture, the used is clean architecture with dependency injection.<p>
<p>
  This architecture has 4 layers
  <ol>
    <li> models</li>
    <li> imagehandler </li>
    <li> middleware </li>
    <li> imagehttphandler </li>
    </ol>
<p>
  
# Model layer
<p> Model layer is reposable to create a structs for our models the to be transfered to the database or bind the requet into it.<br/>
<br />
  
# imagehandler
<p>This contains the repositoris as well as the business logig, this divided into two folders, one called usecase that will handle all the busines logic, and the last one is called repository to handle the database operation </p>

# middleware
<p> This is to handle the crossorigin on the http request </p>

# imagehttphandler
<p> This is responsable to create the http request, this would take the a list of URLS then start to screenshot the urls and the image to the folder path, that would be add to the config.json file </P>

<p><u><b>In this case I am using three packages</b></u></p>
<ol>
   <li> go get github.com/labstack/echo </li>
  <li> go get -u github.com/spf13/viper </li>
  <li> go get github.com/pdfcrowd/pdfcrowd-go</li>
  </ol>
# echo
<p> why echo, because this is provide a high performance, extensible, minimalist for a go web framework/p>
  <p> for more information please check this url : https://echo.labstack.com/ </p>
  
 # viber
 <p> This helo to read the value form the configuration file for more info please check this url : https://github.com/spf13/viper </p>
 
 # pdfcrowd
 <p> This API will help us to take the URL and convert it to imagem and unfortunately this is not free I am using a trial </p>

<p> After the screen have been taken, I moved it to a shared path folder with a UUID to create a unique name for the screenshot image, and then the Image data will be saved to the database, such as folder path and the creation time to the database.</p>
  
