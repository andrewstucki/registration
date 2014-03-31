module.exports = (grunt) ->

  # show elapsed time at the end
  require("time-grunt") grunt

  # load all grunt tasks
  require("load-grunt-tasks") grunt

  siteConfig =
    source: "src"
    test: "test"
    build: ".tmp"
    output: "../server/assets"
    backgrounds: "../server/backgrounds"

  grunt.initConfig
    site: siteConfig
    bower: grunt.file.readJSON('.bowerrc'),
    watch:
      templates:
        files: "<%= site.source %>/templates/**/*.jade"
        tasks: ["jade:templates"]

      coffee:
        files: ["<%= site.source %>/coffee/**/*.coffee"]
        tasks: ["coffee:build"]

      coffeeTest:
        files: ["<%= site.test %>/spec/{,*/}*.coffee"]
        tasks: ["coffee:test"]

      compass:
        files: ["<%= site.source %>/sass/**/*.{scss,sass}"]
        tasks: ["compass:build"]

      scripts:
        files: ["<%= site.build %>/scripts/**/*.js", "!<%= site.build %>/scripts/combined-application.js"]
        tasks: ["scripts:browserify", "scripts:concat"]

    clean:
      build:
        options:
          force: true
        files: [
          dot: true
          src: ["<%= site.build %>", "<%= site.output %>/*", "<%= site.backgrounds %>/*", "!<%= site.output %>/.git*"]
        ]

    coffee:
      build:
        files: [
          expand: true
          cwd: "<%= site.source %>/coffee"
          src: "**/*.coffee"
          dest: "<%= site.build %>/scripts"
          ext: ".js"
        ]

      test:
        files: [
          expand: true
          cwd: "<%= site.test %>/spec"
          src: "**/*.coffee"
          dest: "<%= site.build %>/spec"
          ext: ".js"
        ]

    compass:
      options:
        sassDir: "<%= site.source %>/scss"
        specify: "<%= site.source %>/scss/main.scss"
        cssDir: "<%= site.output %>"
        imagesDir: "<%= site.source %>/images"
        javascriptsDir: "<%= site.source %>/js"
        fontsDir: "<%= site.source %>/font"
        importPath: "<%= bower.directory %>"
        httpImagesPath: "images"
        httpGeneratedImagesPath: "images/generated"
        httpFontsPath: "font"
        relativeAssets: false
        outputStyle: "compressed"
        noLineComments: true
        environment: "production"

      build: {}
      server:
        options:
          debugInfo: false

    # Put files not handled in other tasks here
    copy:
      build:
        files: [
          expand: true
          dot: true
          cwd: "<%= site.source %>"
          dest: "<%= site.output %>"
          src: ["*.ico", "font/*", "images/*"]
        ]
      backgrounds:
        files: [
          expand: true
          dot: true
          cwd: "<%= site.source %>/backgrounds"
          dest: "<%= site.backgrounds %>"
          src: ["*.jpg"]
        ]

    uglify:
      scripts:
        files:
          "<%= site.output %>/application.js": ["<%= site.output %>/application.js"]

    jade:
      templates:
        files: [
          expand: true
          cwd: "<%= site.source %>/templates"
          src: ["**/*.jade"]
          dest: "<%= site.output %>"
          ext: ".html"
        ]

    handlebars:
      compile:
        options:
          namespace: false
          commonjs: true
          processName: (filePath) ->
            filePath = filePath.split(/[\/]+/).pop()
            return filePath.split(/\./)[0]
        files:
          '<%= site.build %>/scripts/templates.js': [
            '<%= site.source %>/coffee/partials/*.hbs'
          ]

    concurrent:
      options:
        limit: 5
      build: ["jade", "coffee", "handlebars", "compass:build", "copy"]
      test: ["jade", "coffee", "handlebars", "compass"]

    karma:
      unit:
        configFile: "karma.conf.js"

    browserify:
      build:
        src: ["<%= site.build %>/scripts/application.js"]
        dest: "<%= site.build %>/scripts/combined-application.js"
      options:
        alias: []

    concat:
      scripts:
        src: [
          "<%= bower.directory %>/modernizr/modernizr.js"
          "<%= bower.directory %>/jquery/jquery.js"
          "<%= bower.directory %>/underscore/underscore.js"
          "<%= bower.directory %>/backbone/backbone.js"
          "<%= bower.directory %>/handlebars/handlebars.js"
          "<%= site.source %>/vendor/jquery.maskedinput.js"
          "<%= site.source %>/vendor/chosen.jquery.js"
          "<%= site.source %>/vendor/jquery.validator.js"
          "<%= site.source %>/vendor/additional-methods.js"
          "<%= bower.directory %>/sass-bootstrap/js/bootstrap-alert.js"
          "<%= bower.directory %>/sass-bootstrap/js/bootstrap-transition.js"
          "<%= bower.directory %>/sass-bootstrap/js/bootstrap-button.js"
          "<%= site.build %>/scripts/combined-application.js"
        ]
        dest: "<%= site.output %>/application.js"

  grunt.registerTask "test", [
    "clean",
    "concurrent:test"
    # "mocha"
  ]

  grunt.registerTask "build", [
    "clean:build",
    "concurrent:build",

    "browserify",
    "concat",
    "uglify",
  ]

  grunt.registerTask "default", ["build"]