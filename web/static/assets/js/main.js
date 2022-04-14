$(document).ready(function(){

    // $('body').addClass('modal-open')
    $(window).on("load", function () {
      $('.loading-page__logo').fadeOut();
      $('.loading-page').delay(350).fadeOut('slow');
      $('body').removeClass('modal-open')
    })

    // Stick header
    var stick_header = $('.header.-fix')
    if( stick_header.length ){
        if ( stick_header.offset().top >= 100 ) stick_header.addClass("is-sticky");
        $(window).scroll(function(){
            $(this).scrollTop()>10?stick_header.addClass("is-sticky"):stick_header.removeClass("is-sticky")
        })
    }

    /*----Get Header Height ---*/
    function get_header_height() {
        var header_sticky = $("header").outerHeight()
        $('body').css("--header-height",header_sticky+'px')
    }

    $(window).on('load resize', function () {
        setTimeout(function(){
            get_header_height()
        }, 500);
    });

    if($('body').hasClass( "home" )){
        new WOW().init();
    }

    //-------------------------------------------------
    // Menu
    //-------------------------------------------------
    $.fn.dnmenu = function( options ) {

        let thiz = this
        let menu = $(this).attr('data-id')
        let menu_id = '#'+menu

        // Default options
        var settings = $.extend({
            name: 'Menu'
        }, options );

        // get ScrollBar Width
        function getScrollBarWidth () {
            var $outer = $('<div>').css({visibility: 'hidden', width: 100, overflow: 'scroll'}).appendTo('body'),
                widthWithScroll = $('<div>').css({width: '100%'}).appendTo($outer).outerWidth();
            $outer.remove();
            return 100 - widthWithScroll;
        };
        let ScrollBarWidth = getScrollBarWidth() + 'px';

        // Create wrap
        // Button click
        thiz.click(function(e){

            e.preventDefault()

            if($(this).hasClass('active')){
                $(this).removeClass('active')
                $('body').removeClass('modal-open')
                $(menu_id).removeClass('active')
                thiz.removeClass('active')
                $('.header-mb').removeClass('-menu-mb-active')

            } else {
                $(this).addClass('active')
                $('body').addClass('modal-open')
                $(menu_id).addClass('active')
                thiz.addClass('active')
                $('.header-mb').addClass('-menu-mb-active')

            }
        });

        // Custom close
        $('.js-menu__close').click(function(){
            $('body').removeClass('modal-open')
            $(thiz).removeClass('active')
            $(menu_id).removeClass('active')
        })

        // Menu
        var el= $(menu_id).find(".nav__mobile--ul");
        el.find(".menu-item-has-children>a").after('<button class="nav__mobile__btn"><i></i></button>'),

        el.find(".nav__mobile__btn").on("click",function(e){
            e.stopPropagation(),
            $(this).parent().find('.sub-menu').first().is(":visible")?$(this).parent().removeClass("sub-active"):
            $(this).parent().addClass("sub-active"),
            $(this).parent().find('.sub-menu').first().slideToggle()
        })

        $('.nav__mobile, .header, .header-mb').mousedown(function(e){ e.stopPropagation(); });
        $(document).mousedown(function(e){ $('.nav__mobile').removeClass('active'); $(thiz).removeClass('active'); $("body").removeClass('modal-open') });

        // Apply options
        return;
    };

    $('.menu-mb__btn').dnmenu()

    //Home page
    var $carousel = $('.nav-roadmap')
    $carousel.on( 'change.flickity', function( event, index ) {
      console.log( 'Slide changed to ' + index )
      var index2 = index + 1
      console.log($('.nav-roadmap li:nth-child('+index2+') .nav-link').html())
      $('.nav-roadmap li:nth-child('+index2+') .nav-link').click()
    });

    // End home page


    // check checkbox checked enable button
    function checkDisableButton() {
        $(".js-attr-disable").each(function () {
          let checkBox = $(this).find(".js-attr-checkbox");
          let btn = $(this).find(".js-attr-button");
          checkBox.on("click", function () {
            console.log(1)
            if ($(this).is(":checked")) {
              $(btn).removeAttr("disabled").removeClass("-disabled");
            } else {
              $(btn).attr("disabled", "disabled").addClass("-disabled");
            }
          });
        });
    }
    checkDisableButton();

    /*Page vesting*/
    $('.js-tr-toggle .td-toggle').click(function(e){
        $(this).parent().next().toggle()
    })

    /*PAge market*/
    $('.js-widget--toggle').on("click",function(e) {
        $('.market__sidebar').toggleClass('active')
        $('body').toggleClass('filter-open')
    })

    $('.js-checkall').on("click",function(e) {
      let checkbox = $(this).closest('ul').find('input[type="checkbox"]:not(.js-checkall)')
      let stt = false
      if($(this).is(":checked") == true) stt = true;
      checkbox.each(function() { this.checked = stt; });
    })

    $('.js-checkbox-list').change(function(e) {
      let parent_checkall = $(this).closest('.js-checkbox-list').find('.js-checkall')
      let checkbox = $(this).closest('ul').find('input[type="checkbox"]:not(.js-checkall)')
      checkbox.each(function() {
        if(this.checked == false){
          parent_checkall.prop('checked', false)
        }
      });
    })
    /*end checkbox*/

    /*PAge Wallet*/
    // Tooltip page account
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
      return new bootstrap.Tooltip(tooltipTriggerEl)
    })
});


