#!/usr/bin/env python

import urllib2
import json
import os

urls = {'classes' : [ 'classes',
                      'SubClasses',
                      'features',
                      'spellcasting',
                      'startingequipment' ],
        'races' : [ 'races',
                    'subraces',
                    'traits' ],
        'equipement' : [ 'equipment',
                         'weapon-properties' ],
        'game_mechanic' : [ 'conditions',
                            'damage-types',
                            'magic-schools' ],
        'monster': [ 'monsters' ],
        'spells' : [ 'spells' ],
        'character_data': [ 'ability-scores',
                            'skills',
                            'proficiencies',
                            'languages']
}

def get_data(url):
    if url.startswith('http://www.dnd5eapi.co/api/'):
        request = urllib2.urlopen(url)
    else:
        request = urllib2.urlopen('http://www.dnd5eapi.co/api/'+url)
    data = json.load(request)   
    return data

for k, v in urls.items():
    if not os.path.exists(k):
        os.makedirs(k)
    
    for i in v:
        
        j = 0
        for items in get_data(i)['results']:
            jdata = get_data(items['url'])

            filename = os.path.join(k,i+'_'+str(j)+'.json')
            print filename
            j = 1+j
            
            if not os.path.exists( filename ) :
                with open( filename, 'w') as datafile:
                    json.dump(jdata, datafile)

            
        
