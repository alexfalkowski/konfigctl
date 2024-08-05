# frozen_string_literal: true

When('I download the configuration') do
  env = {
    'KONFIG_CONFIG_FILE' => '.config/client.yaml',
    'KONFIG_APP_CONFIG_FILE' => 'reports/server.yaml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'config')
  pid = spawn(env, cmd, %i[out err] => ['reports/client.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a configuration') do
  expect(File.file?('reports/server.yaml')).to be true
  expect(@status.exitstatus).to eq(0)
end
