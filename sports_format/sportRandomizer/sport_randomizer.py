import sys
import random
import math 
import json 

def generateSport(skill_count,ar_rules,con_rules,eg_rules,ig_rules):
    sport = {"discipline": "Random sport"}
    skills_array = []
    skills = []
    for i in range(0,skill_count):
        val = random.random()*10
        ded = -random.random()
        skills_array.append(
            {
                "name": "Skill"+str(i),
                "label": "S"+str(i),
                "value": math.floor(val * 100)/100.0,
                "deduction": math.floor(ded * 100)/100.0,
            }
        )
        skills.append("S"+str(i))
    ar = []
    ar_skills = skills.copy()
    max_in_ar = math.floor(skill_count/ar_rules)
    for i in range(0,ar_rules):
        ar_rule = {"k":random.randint(1,math.floor(skill_count/3))}
        chosen = random.choices(ar_skills,k=math.floor(random.random()*max_in_ar)+1)
        chosen = list(dict.fromkeys(chosen))
        ar_rule["skills"] = chosen 
        for sk in chosen:
            ar_skills.remove(sk)
        ar.append(ar_rule)

    con = [] 
    for i in range(0,con_rules):
        con_rule = {"value":math.floor(random.random()*10 * 100)/100.0}
        con_rule["s1"] = random.choice(skills)
        con_rule["s2"] = random.choice(skills)
        con.append(con_rule)

    ig = [] 
    for i in range(0,ig_rules):
        ig_rule = {}
        ig_rule["s1"] = random.choice(skills)
        ig_rule["s2"] = random.choice(skills)
        ig.append(ig_rule)

    eg = []
    eg_skills = skills.copy()
    max_in_eg = math.floor(skill_count/eg_rules)
    for i in range(0,eg_rules):
        eg_rule = {"value":random.randint(1,10)}
        chosen = random.choices(eg_skills,k=math.floor(random.random()*max_in_eg)+1)
        chosen = list(dict.fromkeys(chosen))
        eg_rule["skills"] = chosen 
        for sk in chosen:
            eg_skills.remove(sk)
        eg.append(eg_rule)

    sport["skills"] = skills_array
    sport["antiRepetitionRule"] = {"groups":ar}
    sport["elementGroupRule"] = {"groups":eg}
    sport["connectionRule"] = {"connections":con}
    sport["incompleteGraphRule"] = {"edges":ig}
    return sport

def save_sport(sport):
    with open("randomSport.json", "w") as outfile:
        json.dump(sport, outfile)

if __name__ == '__main__':
    skill_count = 40
    ar_rules = 10
    con_rules = 10
    eg_rules = 10
    ig_rules = 10
    sport_config = generateSport(skill_count,ar_rules,con_rules,eg_rules,ig_rules)
    save_sport(sport_config)