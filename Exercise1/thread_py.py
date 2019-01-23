from threading import Thread

i = 0
number = 1000000

def increase():
	global i

	for j in range(number):
		i += 1

def decrease():
	global i

	for j in range(number):
		i -= 1

def main():

	incrementing = Thread(target = increase, args = (),)
    decrementing = Thread(target = decrease, args = (),)
    
    incrementing.start()
    decrementing.start()
    
    incrementing.join()
    decrementing.join()
    

	print i

main()