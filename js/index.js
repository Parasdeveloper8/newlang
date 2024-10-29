const express = require("express");
const app = express();

const getdata = require("./Routes/getdata");

app.use(express.json());
app.use(express.urlencoded({extended:false}));

app.use("/getdata",getdata);

app.listen(7800,()=>{
    console.log("server is listening on port 7800");
})