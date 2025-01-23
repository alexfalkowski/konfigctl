# frozen_string_literal: true

Nonnative.configure do |config|
  config.load_file('nonnative.yml')
end

require 'nonnative/startup'

Before do
  files = [
    'reports/server.yaml',
    'reports/ssm.secret',
    'reports/vault.secret'
  ]
  files.each { |f| FileUtils.rm_f(f) }
end
