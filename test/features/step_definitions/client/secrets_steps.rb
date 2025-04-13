# frozen_string_literal: true

When('I write secrets') do
  env = {
    'KONFIG_CONFIG_FILE' => '.config/client.yaml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'secrets', '-i env:KONFIG_CONFIG_FILE')
  pid = spawn(env, cmd, %i[out err] => ['reports/secrets.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('I try to write missing secrets') do
  env = {
    'KONFIG_CONFIG_FILE' => '.config/invalid.yaml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'secrets', '-i env:KONFIG_CONFIG_FILE')
  pid = spawn(env, cmd, %i[out err] => ['reports/secrets.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have secrets') do
  expect(@status.exitstatus).to eq(0)
  expect(File).to exist('reports/vault.secret')
  expect(File).to exist('reports/ssm.secret')
end

Then('I should not have secrets') do
  expect(@status.exitstatus).to eq(1)
  expect(File).not_to exist('reports/vault.secret')
  expect(File).not_to exist('reports/ssm.secret')
end
