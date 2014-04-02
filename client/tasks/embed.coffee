ResourceEmbedder = require 'resource-embedder'
path = require 'path'

module.exports = (grunt) ->

  grunt.registerMultiTask 'embed', 'Converts external scripts and stylesheets into embedded ones.', ->
    done = @async()
    index = 0
    count = @files.length
    
    grunt.log.ok "Starting embed."
    
    @files.forEach (file) =>
      
      filepath = file.src.toString()
      
      if typeof filepath isnt 'string' or filepath is ''
        grunt.log.error 'src must be a single string'
        return false
      
      if !grunt.file.exists filepath
        grunt.log.error 'Source file "' + filepath + '" not found.'
        return false
            
      embedder = new ResourceEmbedder filepath, @options()
      embedder.get (output, warnings) ->
        grunt.file.write file.dest, output
        if warnings?
          grunt.log.warn warning for warning in warnings
        grunt.log.ok "File \"#{file.dest}\" created."
        index++
        if index is count
          done()
      