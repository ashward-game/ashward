import datetime
from datetime import timezone
from dateutil import relativedelta


dtFormat = "%Y-%m-%d %H:%M:%S%Z"
startDate = "18/04/22 16:30:00"
gameDate = "25/03/22 16:30:00"
deployDate = "18/03/22 16:30:00"
startDT = datetime.datetime.strptime(startDate, '%d/%m/%y %H:%M:%S')
startDT = startDT.replace(tzinfo=timezone.utc)
gameDT = datetime.datetime.strptime(gameDate, '%d/%m/%y %H:%M:%S')
gameDT = gameDT.replace(tzinfo=timezone.utc)
deployDT = datetime.datetime.strptime(deployDate, '%d/%m/%y %H:%M:%S')
deployDT = deployDT.replace(tzinfo=timezone.utc)


def genTGEpercent(num):
    return '''_tge_percent = {0}; // {1:.2f}%\n'''.format(int(num), num/100)


def deployTimestamp():
    return str(int(deployDT.timestamp()))


def genClaimableMilestones(cliff, vest):
    dt = startDT
    solCode = "_claimableMilestones = ["
    solCode += deployTimestamp() + ", "
    for i in range(cliff):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(vest-cliff):
        solCode += str(int(dt.timestamp()))
        if i < vest-cliff - 1:
            solCode += ", "
        dt = dt + relativedelta.relativedelta(months=1)
    solCode += "];\n"
    return solCode


def genClaimableMilestonesSpecial(cliff1, vest1, cliff2, vest2):
    dt = startDT
    solCode = "_claimableMilestones = ["
    solCode += deployTimestamp() + ", "
    for i in range(cliff1):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(vest1-cliff1):
        solCode += str(int(dt.timestamp()))
        solCode += ", "
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(cliff2):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(vest2-cliff2):
        solCode += str(int(dt.timestamp()))
        if i < vest2-cliff2 - 1:
            solCode += ", "
        dt = dt + relativedelta.relativedelta(months=1)
    solCode += "];\n"
    return solCode


def genClaimableMilestonesGame(cliff, vest):
    dt = startDT
    solCode = "_claimableMilestones = ["
    solCode += deployTimestamp() + ", "
    solCode += str(int(gameDT.timestamp())) + ", "
    for i in range(cliff):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(vest-cliff):
        solCode += str(int(dt.timestamp()))
        if i < vest-cliff - 1:
            solCode += ", "
        dt = dt + relativedelta.relativedelta(months=1)
    solCode += "];\n"
    return solCode


def genClaimablePercent(stamp, num, dt):
    return '''_claimablePercents[{}] = {}; // {}\n'''.format(stamp, int(num), datetime.datetime.strftime(dt, dtFormat))


def genClaimablePercents(tge, num, cliff, vest):
    realVest = vest - cliff
    dt = startDT
    solCode = ""
    for i in range(cliff):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(realVest - 1):
        solCode += genClaimablePercent(int(dt.timestamp()), num, dt)
        dt = dt + relativedelta.relativedelta(months=1)
    lastPercent = 10000 - (realVest - 1) * num - tge
    solCode += genClaimablePercent(int(dt.timestamp()), lastPercent, dt)
    return solCode


def genClaimablePercentsGame(tge, num, cliff, vest):
    realVest = vest - cliff
    dt = startDT
    solCode = ""
    solCode += genClaimablePercent(int(gameDT.timestamp()), num, gameDT)
    for i in range(cliff):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(realVest - 1):
        solCode += genClaimablePercent(int(dt.timestamp()), num, dt)
        dt = dt + relativedelta.relativedelta(months=1)
    lastPercent = 10000 - (realVest - 1) * num - tge
    solCode += genClaimablePercent(int(dt.timestamp()), lastPercent, dt)
    return solCode


def genClaimablePercentsSpecialStart(tge, num, cliff, vest):
    realVest = vest - cliff
    dt = startDT
    solCode = ""
    for i in range(cliff):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(realVest):
        solCode += genClaimablePercent(int(dt.timestamp()), num, dt)
        dt = dt + relativedelta.relativedelta(months=1)
    return solCode


def genClaimablePercentsSpecialEnd(tge, start, num, phase1, cliff, vest):
    realVest = vest - cliff
    dt = startDT
    solCode = ""
    for i in range(phase1):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(cliff):
        dt = dt + relativedelta.relativedelta(months=1)
    for i in range(realVest - 1):
        solCode += genClaimablePercent(int(dt.timestamp()), num, dt)
        dt = dt + relativedelta.relativedelta(months=1)
    lastPercent = 10000 - (realVest - 1) * num - tge - start
    solCode += genClaimablePercent(int(dt.timestamp()), lastPercent, dt)
    return solCode


def strategicPartner():
    # TGE: 0%, MUR: 8.33%, 6 months cliff, vest monthly in 18 months
    solCode = genTGEpercent(0)
    solCode += genClaimableMilestones(6, 18)
    solCode += genClaimablePercents(0, 833, 6, 18)
    return solCode


def privateSpecial():
    # note: 18.03 = 1st month
    # TGE 8%. lock 2nd, 3rd. Month 4th vest 6.13%. Lock month 5th. From month 6th vest 6.13% monthly.
    # TGE 8%, MUR 6.13%, phase1: cliff 3 months, vest 4 month, phase 2: cliff 1, vest: 15
    solCode = genTGEpercent(800)
    solCode += genClaimableMilestonesSpecial(2, 3, 1, 15)
    solCode += genClaimablePercentsSpecialStart(800, 613, 2, 3)
    solCode += genClaimablePercentsSpecialEnd(800, 613, 613, 3, 1, 15)
    return solCode


def IDO():
    # TGE 20%, MUR: 20% 1 month cliff, vest monthly in 5 months
    solCode = genTGEpercent(2000)
    solCode += genClaimableMilestones(1, 5)
    solCode += genClaimablePercents(2000, 2000, 1, 5)
    return solCode


def advisory():
    # TGE 6%, MUR: 7.83%, vest monthly in 12 months
    solCode = genTGEpercent(600)
    solCode += genClaimableMilestones(0, 12)
    solCode += genClaimablePercents(600, 783, 0, 12)
    return solCode


def team():
    # TGE: 0%, MUR: 8.33%, 12 months cliff, vest monthly 24 months
    solCode = genTGEpercent(0)
    solCode += genClaimableMilestones(12, 24)
    solCode += genClaimablePercents(0, 833, 12, 24)
    return solCode


def marketing():
    # TGE: 0%, MUR: 4.35%, 1 month cliff, vest monthly in 24 months
    solCode = genTGEpercent(0)
    solCode += genClaimableMilestones(1, 24)
    solCode += genClaimablePercents(0, 435, 1, 24)
    return solCode


def play():
    # TGE: 0%, Start vesting upon the game launch, 2.78% monthly, vest monthly in 36 months
    solCode = genTGEpercent(0)
    solCode += genClaimableMilestonesGame(0, 35)
    solCode += genClaimablePercentsGame(278, 278, 0, 35)
    return solCode


def staking():
    # TGE: 0%, Start vesting upon the game launch, 2.78% monthly, vest monthly in 36 months
    solCode = genTGEpercent(0)
    solCode += genClaimableMilestonesGame(0, 35)
    solCode += genClaimablePercentsGame(278, 278, 0, 35)
    return solCode


def liquidity():
    # TGE 20%, MUR 8.89%, 3 months cliff, vest monthly in 12 months
    solCode = genTGEpercent(2000)
    solCode += genClaimableMilestones(3, 12)
    solCode += genClaimablePercents(2000, 889, 3, 12)
    return solCode


def reserve():
    # TGE: 0%, MUR: 2.78%, Vest monthly in 36 months
    solCode = genTGEpercent(0)
    solCode += genClaimableMilestones(0, 36)
    solCode += genClaimablePercents(0, 278, 0, 36)
    return solCode


if __name__ == "__main__":
    # print(strategicPartner())
    # print(privateSpecial())
    # print(IDO())
    # print(advisory())
    print(team())
    # print(marketing())
    # print(play())
    # print(staking())
    # print(liquidity())
    # print(reserve())
