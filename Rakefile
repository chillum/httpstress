# A script to cross-compile Go project and ZIP the binaries.
# Settings are specified in a YAML file: build.yml.
#
# Example config:
#
# platforms:
#   - name: linux
#     arch:
#       - 386
#       - amd64
#   - name: windows
#     arch: amd64
#   - name: windows
#     arch: 386
# out: ~/Downloads
#
# Please note that if 'out' is specified, ZIP files will appear in the specified directory;
# if not, they will be in current directory.

# Copyright 2014 Chai Chillum
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

require 'yaml'
config = YAML.load_file 'build.yml' # Rake fails to run on problems with config file.

desc 'Build this project for the platforms in build.yml'
task :build do
  config['platforms'].each do |os|
    if os['arch'].respond_to?('each')
      os['arch'].each do |arch|
        build os['name'], arch
      end
    else
      build os['name'], os['arch']
    end
  end
end

desc 'Run `go test` for the platforms in build.yml'
task :test do
  config['platforms'].each do |os|
    if os['arch'].respond_to?('each')
      os['arch'].each do |arch|
        test os['name'], arch
      end
    else
      test os['name'], os['arch']
    end
  end
end

desc 'ZIP this project binaries'
task :zip => [:build, :test] do
  unless config['out']
    config['out'] = '.'
  end

  config['platforms'].each do |os|
    if os['arch'].respond_to?('each')
      os['arch'].each do |arch|
        setenv os['name'], arch
        zip "#{config['out']}/#{os['name']}_#{arch}", `go list -f '{{.Target}}'`
      end
    else
        setenv os['name'], os['arch']
        zip "#{config['out']}/#{os['name']}_#{os['arch']}", `go list -f '{{.Target}}'`
    end
  end
end

def setenv os, arch
  ENV['GOARCH'] = arch.to_s
  ENV['GOOS']   = os
end

def build os, arch
  setenv os, arch

  puts "Building #{os}_#{arch}"
  system('go install')
end

def test os, arch
  setenv os, arch

  puts "Testing #{os}_#{arch}"
  if system('go test') != true
    puts 'Tests failed. Exiting'
    exit 1 # Rake returns 1 if tests for some arch fail.
  end
end

def zip target, file
  if system("zip -qj #{target}.zip #{file}") == true
    puts "Wrote #{target}.zip"
  end
end
