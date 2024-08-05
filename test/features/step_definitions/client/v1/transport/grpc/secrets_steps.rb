# frozen_string_literal: true

When('I write secrets') do
  env = {
    'KONFIG_CONFIG_FILE' => '.config/client.yaml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'secrets')
  pid = spawn(env, cmd, %i[out err] => ['reports/client.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have secrets') do
  expect(File).to exist('reports/vault.secret')
  expect(File).to exist('reports/ssm.secret')
end
