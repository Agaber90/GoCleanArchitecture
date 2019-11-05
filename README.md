# detectifyScreenshot
This a Golang project to start take a screenshot form URLS using pdfcrowd API and save the images data to mysql database
before start to explain this porject please

# install the following packages 
<br />
<ol>
  
  <li> go get github.com/go-sql-driver/mysql </li>
  <li> go get github.com/labstack/echo </li>
  <li> go get -u github.com/spf13/viper </li>
  </ol>
<br>

<p>First of all let me speak about the solution architecture, the used is clean architecture with dependency injection.<p>
<p>
  This architecture has 4 layers
  <ol>
    <li> models</li>
    <li>imagehandler</li>
    <li>middleware</li>
    <li>imagehttphandler</li>
    </ol>
<p>
