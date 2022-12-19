## Rewrite about Cloud compute project 

### overview 

>Note:I have not used a database therfore I used in memory variables instead.

First of all ,I made a `gin Engine `and assigned it to a variable called` router `
Also, I made the Engine to load my basic html files, so for that I've used `router`,loadHTMLGlob("templates/*")
Router variable is an instance of gin Engine ,therefore, the request will be proccessed seperately.
I used a framework named gin ,I created an instance from it ,and saved it with a variable named `Router`
I used routing with gin framework and created a function to handling each route ,
I did not  use the database , I used  in memory variables instead 
Create a second variable for the cart to add `product` to a user who added it to `the card`
#### I used five routes
the first 	router `handleAllProducts` Show HTML page with all products 
the second `handleProduct` Add product to cart 
the third `handleCart` Show HTML page with all products in the cart 
the foruth 	`handleRemoveAll`Remove all products from the cart 
the fifth   handleRemoveById Remove a productes from the cart by its ID
I used `server` from library HTTP Build-in with `GO` language .
The server  listening  on port `8000`

