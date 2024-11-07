module main

import os
import rand
import term
import time

fn main() {
	mut menanginput := 0
	mut menangcpu := 0
	mut i := 0
	kelemahan := {
		'gunting': 'batu'
		'kertas': 'gunting'
		'batu': 'kertas'
	}
	for {
		i++
		term.clear()
		komp := go rand.element(kelemahan.values())
		println('Ronde $i: Komputer $menangcpu Anda $menanginput')
		println('Gunting, batu atau kertas!')
		input := os.input('Pilih seranganmu! ').to_lower()
		cpu := komp.wait() or { panic('$err') }
		println('Serangan mu: $input Serangan komputer: $cpu')
		match true {
			input == kelemahan[cpu] { 
				println(term.green('Kamu menang!')) 
				menanginput++
			}
			kelemahan[input] == cpu { 
				println(term.red(r'Yah... kamu kalah. :(')) 
				menangcpu++
			}
			input == cpu { println('Kamu seri.') }
			else { println(term.red('Input invalid!')) }
		}
		time.sleep(1500 * time.millisecond)
	}
}
