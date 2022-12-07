// $(document).ready(function(){
//     $('#dropdownlogin').click(function(event){
//         // your stuff here
//         alert("Alert On Click");
//         event.stopPropagation();
//     });
// });
// $('.dropdown-menu').on("click.bs.dropdown", function (e) { 
//                  e.stopPropagation();        
//                  e.preventDefault();    
//                  console.log("F")           
// });
// $('#dropdown').click(function(event) {
    
//     $(this).parent().toggleClass('open');
// });
//$('#dropdownlogin','#dropdown').on('hide.bs.dropdown', function (e){
// $('#dropdownlogin','#dropdown').on('click', function (e){
//     var target = $(e.target);
//     console.log("F") 
//     if(target.hasClass("keepopen") || target.parents(".keepopen").length){
//         return false; // returning false should stop the dropdown from hiding.
//     }else{
//         return true;
//     }
// });

// $('#dropdownlogin','#dropdownMenuClickable').on('hide.bs.dropdown', function(event){
//     $(this).parent().is(".open") && event.stopPropagation()
//     var events = $._data(document, 'events') || {};
//     events = events.click || [];
//     console.log("F") 
//     for(var i = 0; i < events.length; i++) {
//         if(events[i].selector) {

//             //Check if the clicked element matches the event selector
//             if($(event.target).is(events[i].selector)) {
//                 events[i].handler.call(event.target, event);
//             }

//             // Check if any of the clicked element parents matches the 
//             // delegated event selector (Emulating propagation)
//             $(event.target).parents(events[i].selector).each(function(){
//                 events[i].handler.call(this, event);
//             });
//             var dropdown = new bootstrap.Dropdown(element, {
//                 popperConfig: function (defaultBsPopperConfig) {
//                   var newPopperConfig = {show=true}
//                   // use defaultBsPopperConfig if needed...
//                   return newPopperConfig
//                 }
//               })
//         }
//     }
//     event.stopPropagation(); //Always stop propagation
// });