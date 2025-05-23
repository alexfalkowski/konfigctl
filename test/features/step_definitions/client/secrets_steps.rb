# frozen_string_literal: true

Before do
  files = [
    'reports/server.yaml',
    'reports/ssm.secret',
    'reports/vault.secret'
  ]
  files.each { |f| FileUtils.rm_f(f) }
end

When('I write {string} secrets') do |name|
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfigctl', 'secrets', "-i file:.config/#{name}.yaml")
  pid = spawn(cmd, %i[out err] => ['reports/secrets.log', 'a'])

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
