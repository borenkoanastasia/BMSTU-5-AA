import faker

N = 2000
def generate_dict():
    arr = []
    x = "234"
    y = "5678"
    s = ""
    for j in range(5):
        for k in range(10):
            for b in range(10):
                for l in range(10):
                    s = str(j)+str(k)+str(b)+str(l)
                    arr.append(x+s+y)
                    s = ""
    return arr

def print_in_file(arr):
    f = open('dict.txt','w')
    for i in arr:
        if len((i[1].split()))>2:
            continue
        f.write(str(i[0]))
        f.write(" ")
        f.write(str(i[1]))
        f.write("\n")
    f.close()

ARR = generate_dict()
for i in generate_dict():
    print(i, end=", ")
   
RES = []    
        
fake = faker.Faker()
fake.name()

for i in ARR:
    RES.append([i, fake.name()])#.upper()])

print(RES)
print_in_file(RES)
