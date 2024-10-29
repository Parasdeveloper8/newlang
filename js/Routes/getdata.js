const express = require("express");

const router = express.Router();

router.get("/",(req,res)=>{
    res.json({"message":"hello from paras express backend"})
});

module.exports = router;