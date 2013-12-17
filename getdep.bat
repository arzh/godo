@echo off
SET curdir=%CD%
cd ..
git clone https://github.com/arzh/go-clu.git clu
git clone https://github.com/arzh/osext.git
cd %curdir%