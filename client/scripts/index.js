$(document).ready(function() {
  var currentWidth     = String($(window).width())
  var currentHeight    = String($(window).height())
  var firstKeyStrokeAt = null
  var sessionId        = null
  var apiBaseURL       = "http://localhost:3000"

  var postRequest = function(url, data) {
    //takes url and payload and performs a post request
    data["websiteUrl"] = window.location.origin
    data["sessionId"] = sessionId
    
    // returns deferred response
    return $.ajax({
      method: "POST",
      url: url,
      data: JSON.stringify(data)
    })
  }
  
  // initialize all events listeners
  var initListeners = function() {
    
    // to handle both cases if first input was pasted or typed.
    $(".form-control").on('input', function(e) {
      $(".form-control").off("input")
      firstKeyStrokeAt = new Date()
    })
   
    // return a function that when called notifies the API that a copyPasteEvent occured
    var copyPasteHandler = function(pasted) {
      return function(event) {
        postRequest(apiBaseURL, {
          eventType: "copyAndPaste",
          pasted: pasted,
          formId: event.target.id
        })
      }
    }
    
    // register copy & paste events
    $(".form-control").on('copy', copyPasteHandler(false))
    $(".form-control").on('paste', copyPasteHandler(true))
    

    // register resize Event
    $(window).on('resize', function() {
      var newWidth = String($(window).width())
      var newHeight = String($(window).height())
      
      // notify API that resizeEvent occured
      postRequest(apiBaseURL, {
        eventType: "resize",
        resizeFrom: {width: currentWidth, height: currentHeight},
        resizeTo: {width: newWidth, height: newHeight}
      })
      
      // update currentSize
      currentWidth = newWidth
      currentHeight = newHeight
    })
    
    // register for submit
    $(".form-details").submit(function(e) {
      postRequest(apiBaseURL, {
        eventType: "timeTaken",
        time: ((new Date).getTime() - firstKeyStrokeAt.getTime()) / 1000 //originally in milliseconds
      })
    })
  }
  
  // if sessionId is not found in our cookies
  if (!Cookies('sessionId')) {
    // request a new sessionId from the API
    postRequest(apiBaseURL + "/session", {})
      .done(function(data) {
        // then continue initialization
        sessionId = data["sessionId"]
        Cookies.set('sessionId', sessionId)
        initListeners()
    })
  } else {
    //sessionId present. Safe to initailize
    sessionId = Cookies('sessionId')
    initListeners()
  }
})


