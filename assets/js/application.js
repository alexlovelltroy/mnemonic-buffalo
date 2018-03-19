require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");
$(() => {
    var n_match  = ntc.name(hexcolor);
    var n_rgb        = n_match[0]; // This is the RGB value of the closest matching color
    var n_name       = n_match[1]; // This is the text string for the name of the match
    var n_exactmatch = n_match[2]; // True if exact color match, False if close-match

    // alert(n_name);
    $( "span.hexcolor" ).html( n_name );
});
