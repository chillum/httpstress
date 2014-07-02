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
  config['platforms'].each { |os|
    if os['arch'].respond_to?('each')
      os['arch'].each { |arch| build os['name'], arch }
    else
      build os['name'], os['arch']
    end
  }
end

desc 'Run `go test` for the platforms in build.yml'
task :test do
  config['platforms'].each { |os|
    if os['arch'].respond_to?('each')
      os['arch'].each { |arch| test os['name'], arch }
    else
      test os['name'], os['arch']
    end
  }
end

desc 'ZIP this project binaries'
task :zip => [:build, :test] do
  unless config['out']
    config['out'] = '.' # Default to the current directory, if 'out' is not specified.
  end

  config['platforms'].each { |os|
    if os['arch'].respond_to?('each')
      os['arch'].each { |arch| zip os['name'], arch, config['out'] }
    else
      zip os['name'], os['arch'], config['out']
    end
  }
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
  unless system('go test')
    puts 'Tests failed. Exiting'
    exit 1 # Rake returns 1 if tests for some arch fail.
  end
end

def zip os, arch, dir
  setenv os, arch

  if system("zip -qj #{dir}/#{os}_#{arch}.zip #{`go list -f '{{.Target}}'`}")
    puts "Wrote #{dir}/#{os}_#{arch}.zip"
  end
end
