"use strict"
import $ from "jquery"

$(document).ready(function() {

  $("#login_btn").click(()=>{
    alert($("username").val());
  })

})
