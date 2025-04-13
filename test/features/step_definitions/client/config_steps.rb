# frozen_string_literal: true

When('I download the configuration') do
  env = {
    'KONFIG_CONFIG_FILE' => '.config/client.yaml',
    'KONFIG_APP_CONFIG_FILE' => 'reports/server.yaml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'config', '-i env:KONFIG_CONFIG_FILE', '-o env:KONFIG_APP_CONFIG_FILE')
  pid = spawn(env, cmd, %i[out err] => ['reports/config.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('I download a missing configuration') do
  env = {
    'KONFIG_CONFIG_FILE' => '.config/invalid.yaml',
    'KONFIG_APP_CONFIG_FILE' => 'reports/server.yaml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'config', '-i env:KONFIG_CONFIG_FILE', '-o env:KONFIG_APP_CONFIG_FILE')
  pid = spawn(env, cmd, %i[out err] => ['reports/config.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a configuration') do
  expect(@status.exitstatus).to eq(0)
  expect(File).to exist('reports/server.yaml')
end

Then('I should not have a configuration') do
  expect(@status.exitstatus).to eq(1)
  expect(File).not_to exist('reports/server.yaml')
end
