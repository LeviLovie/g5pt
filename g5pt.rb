# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class G5pt < Formula
  desc "App for encrypt and decrypt files"
  homepage "https://github.com/LeviiLovie/g5pt"
  url "https://github.com/LeviiLovie/g5pt/archive/refs/tags/1.0.0.tar.gz"
  sha256 "ca10f5763f0dd616223b08a5d5df62172ebe7237e8f82478b1c11c039cebe4f6"
  license "MIT"

  # depends_on "cmake" => :build

  def install
    bin.install "g5pt"
    bin.install Dir["lib"]
    bin.install Dir["files]
    prefix.install "README.md"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test g5pt`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
