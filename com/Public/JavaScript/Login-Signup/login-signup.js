$(document).ready(function(){
    $(".login-sigup-container .rgstr-btn button").click(function(){
        $('.login-sigup-container .wrapper').addClass('move');
        $('.body-login-signup').css('background','#292c2f');
        $(".login-sigup-container .login-btn button").removeClass('active');
        $(this).addClass('active');

    });
    $(".login-sigup-container .login-btn button").click(function(){
        $('.login-sigup-container .wrapper').removeClass('move');
        $('.body-login-signup').css('background','#ff4931');
        $(".login-sigup-container .rgstr-btn button").removeClass('active');
        $(this).addClass('active');
    });
});