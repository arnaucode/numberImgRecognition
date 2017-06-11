from PIL import Image
import numpy as np
import matplotlib.pyplot as plt
import time
from collections import Counter
import operator

def createExamples():
    numberArrayExamples = open('numArEx.txt', 'a')
    numbersWeHave = range(1,10)
    versionsWeHave = range(1,10)

    for eachNum in numbersWeHave:
        for eachVer in versionsWeHave:
            #print str(eachNum) + '.' + str(eachVer)
            imgFilePath = 'images/numbers/' + str(eachNum) + '.' + str(eachVer) + '.png'
            ei = Image.open(imgFilePath)
            eiar = np.array(ei)
            eiar1 = str(eiar.tolist())

            lineToWrite = str(eachNum) + '::' + eiar1 + '\n'
            numberArrayExamples.write(lineToWrite)


def threshold(imageArray):
    balanceAr = []
    newAr = imageArray

    for eachRow in imageArray:
        for eachPix in eachRow:
            avgNum = reduce(lambda x, y: x+y, eachPix[:3])/len(eachPix[:3])
            balanceAr.append(avgNum)

    balance = reduce(lambda x, y: x + y, balanceAr)/len(balanceAr)
    
    for eachRow in newAr:
        for eachPix in eachRow:
            if reduce(lambda x, y: x+y, eachPix[:3])/len(eachPix[:3]) > balance:
                eachPix[0] = 255
                eachPix[1] = 255
                eachPix[2] = 255
                eachPix[3] = 255
            else:
                eachPix[0] = 0
                eachPix[1] = 0
                eachPix[2] = 0
                eachPix[3] = 255
    
    return newAr

def whatNumIsThis(filePath):
    matchedAr = []
    loadExamps = open('numArEx.txt', 'r').read()
    loadExamps = loadExamps.split('\n')

    i = Image.open(filePath)
    iar = np.array(i)
    iarl = iar.tolist()

    inQuestion = str(iarl)
    
    for eachExample in loadExamps:
        if len(eachExample) > 3:
            splitEx = eachExample.split('::')
            currentNum = splitEx[0]
            currentAr = splitEx[1]

            eachPixEx = currentAr.split('],')

            eachPixInQ = inQuestion.split('],')

            x=0

            while x < len(eachPixEx):
                if eachPixEx[x] == eachPixInQ[x]:
                    matchedAr.append(int(currentNum))

                x += 1
    
    #print matchedAr
    x = Counter(matchedAr)
    #print x
    numb =  max(x.iteritems(), key=operator.itemgetter(1))[0]
    #print numb
    return numb, x 



n, counts = whatNumIsThis('images/test.png')
print n
print counts
