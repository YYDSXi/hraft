
function ReflushTime() {
    #获取当前时间 后一分钟时间
    #DATE=$(date -d '-1 minute ago' +'%Y-%m-%d %H:%M:%S')

    ymdNow=$(date +'%Y-%m-%d')
    hmsNow=$(date -d '0 minute ago' +'%H:%M:%S')
    ymd_hms_Now=$ymdNow" "$hmsNow
    globalTime123_Now=$ymd_hms_Now".123"
    globalTime456_Now=$ymd_hms_Now".456"

    #1多笔无异常存证数据
    dataReceipt1_1T="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"r21r4f431feqf\",\"KeyId\":\"f491gf91\",\"ReceiptValue\":10.1,\"Version\":\"v1.0\",\"UserName\":\"ls\",\"OperationType\":\"111\",\"DataType\":\"test\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"2r45fr\",\"AttachmentFileUris\":[\"fewqf\",\"g52gtttg\"],\"AttachmentTotalHash\":\"3f41f431\"}"
    dataReceipt1_2T="{\"CreateTimestamp\":\""${globalTime456_Now}"\",\"EntityId\":\"f489f11347gf01g\",\"KeyId\":\"0gbw979g\",\"ReceiptValue\":14.5,\"Version\":\"v2.0\",\"UserName\":\"rggsd\",\"OperationType\":\"54gd\",\"DataType\":\"452g5\",\"ServiceType\":\"g35gwe\",\"FileName\":\"测试数据2.txt\",\"FileSize\":432.20,\"FileHash\":\"13g51gegw51g46ragwwg2g5\",\"Uri\":\"/usr/local/2.txt\",\"ParentKeyId\":\"15f1gdr\",\"AttachmentFileUris\":[\"fewqf\",\"g52gfwg\"],\"AttachmentTotalHash\":\"9fdg94f1\"}"

    #2多笔无异常交易数据
    transaction2_1T="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"fg79v9fg7rr1g\",\"TransactionId\":\"07dnsjhja\",\"Initiator\":\"g52\",\"Receipt\":\"108h3801\",\"TxAmount\":103.3,\"DataType\":\"测试123001\",\"ServiceType\":\"e218hen\",\"Remark\":\"312080hr83\",\"BlockIdentify\":\"219e\"}"
    transaction2_2T="{\"CreateTimestamp\":\""${globalTime456_Now}"\",\"EntityId\":\"9y742bf0g8qeg\",\"TransactionId\":\"gydsiffsa\",\"Initiator\":\"zjl\",\"Receipt\":\"12djjuyf\",\"TxAmount\":141.5,\"DataType\":\"测试123002\",\"ServiceType\":\"5greoiw\",\"Remark\":\"fd4duigdgh\",\"BlockIdentify\":\"52j4\"}"

    #3多笔时间戳相同存证数据
    dataReceipt3_1T="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"4g01gh1051g\",\"KeyId\":\"2g552gwerg\",\"ReceiptValue\":145.1,\"Version\":\"v1.0\",\"UserName\":\"452g4\",\"OperationType\":\"gw\",\"DataType\":\"test\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"g52g5rewg\",\"AttachmentFileUris\":[\"g52g5g\",\"431\"],\"AttachmentTotalHash\":\"g432g52\"}"
    dataReceipt3_2T="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"g52gweg42g\",\"KeyId\":\"2g54g54gg\",\"ReceiptValue\":164.5,\"Version\":\"v2.0\",\"UserName\":\"rgg245ggsd\",\"OperationType\":\"rwe\",\"DataType\":\"452g5\",\"ServiceType\":\"g35gwe\",\"FileName\":\"测试数据2.txt\",\"FileSize\":432.20,\"FileHash\":\"13g51gegw51g46ragwwg2g5\",\"Uri\":\"/usr/local/2.txt\",\"ParentKeyId\":\"gwg54g\",\"AttachmentFileUris\":[\"wgrwg\",\"greg5\"],\"AttachmentTotalHash\":\"grewgweg\"}"

    #4多笔时间戳相同交易数据
    transaction4_1T="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"f48051g526\",\"TransactionId\":\"f431g5\", \"Initiator\":\"34f3\", \"Receipt\":\"g52gwrt5g2\", \"TxAmount\":103.3,\"DataType\":\"8gfg413ffr\", \"ServiceType\":\"08qfr\", \"Remark\":\"52g4gt\", \"BlockIdentify\":\"243543\" }"
    transaction4_2T="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"08qfg9qf44f41\",\"TransactionId\":\"52gtwg\", \"Initiator\":\"rweg\", \"Receipt\":\"g52gtgtr\",\"TxAmount\":141,\"DataType\":\"f08e9q744\", \"ServiceType\":\"43gt\", \"Remark\":\"rgewgwg\", \"BlockIdentify\":\"rewrew\" }"


    #5多笔延时数据存证交易数据
    #获取当前时间 前3分钟时间
    ymdPre=$(date +'%Y-%m-%d')
    hmsPre=$(date -d '+3 minute ago' +'%H:%M:%S')
    ymd_hms_Pre=$ymdPre" "$hmsPre
    globalTime123_Pre=$ymd_hms_Pre".123"
    globalTime456_Pre=$ymd_hms_Pre".456"
    #延时一分钟数据
    dataReceipt5_1T="{\"CreateTimestamp\":\""${globalTime123_Pre}"\",\"EntityId\":\"41F5315\",\"KeyId\":\"T35t52t52\",\"ReceiptValue\":134.1,\"Version\":\"v1.0\",\"UserName\":\"htehte\",\"OperationType\":\"ndhnh\",\"DataType\":\"test\",\"ServiceType\":\"ndhn\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"h4hrgqr97grq9g0qg0q8g63h\"}"
    dataReceipt5_2T="{\"CreateTimestamp\":\""${globalTime456_Pre}"\",\"EntityId\":\"G52G52G62\",\"KeyId\":\"g52462h2h2\",\"ReceiptValue\":13.1,\"Version\":\"v24.0\",\"UserName\":\"htreht\",\"OperationType\":\"ndn\",\"DataType\":\"3g\",\"ServiceType\":\"ndh\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"0xnfh80fh63h63h6newfh0h8f84efef\"}"

    #获取当前时间 前13分钟时间
    ymdPreF=$(date +'%Y-%m-%d')
    hmsPreF=$(date -d '+13 minute ago' +'%H:%M:%S')
    ymd_hms_PreF=$ymdPreF" "$hmsPreF
    globalTime123_PreF=$ymd_hms_PreF".123"
    #延时一分钟数据
    dataReceipt5_3F="{\"CreateTimestamp\":\""${globalTime123_PreF}"\",\"EntityId\":\"H6HTRHERH\",\"KeyId\":\"FFh63h63h6QF413\",\"ReceiptValue\":36.1,\"Version\":\"v1.0\",\"UserName\":\"y4eh\",\"OperationType\":\"dnhnh\",\"DataType\":\"ndhn\",\"ServiceType\":\"tnest\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"h63h63h6\"}"

    #############################################################################################

    #6多笔存证ID相同数据(根据存证ID/交易ID去重)
    dataReceipt6_1F="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"g13g5\",\"KeyId\":\"1234567890\",\"ReceiptValue\":10.1,\"Version\":\"v1.0\",\"UserName\":\"ls\",\"OperationType\":\"111\",\"DataType\":\"test\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"r4141frwgtew\"}"
    dataReceipt6_2F="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"12f92gqg5g7fbw\",\"KeyId\":\"1234567890\",\"ReceiptValue\":11.1,\"Version\":\"v24.0\",\"UserName\":\"sg\",\"OperationType\":\"3443g\",\"DataType\":\"3g\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"31g51g51g1\"}"

    #7存证/交易数据时间戳或ID缺失(过滤掉)
    dataReceipt7_1F="{\"CreateTimestamp\":\"\",\"EntityId\":\"12f927fbw\",\"KeyId\":\"g51g1g1551\",\"ReceiptValue\":11.1,\"Version\":\"v24.0\",\"UserName\":\"sg\",\"OperationType\":\"3443g\",\"DataType\":\"3g\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"t2hth3hy3h\"}"
    dataReceipt7_2F="{\"CreateTimestamp\":\""${globalTime123_Now}"\",\"EntityId\":\"531g51\",\"KeyId\":\"\",\"ReceiptValue\":11.1,\"Version\":\"v24.0\",\"UserName\":\"sg\",\"OperationType\":\"3443g\",\"DataType\":\"3g\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"grweg5g25g52\"}"


    #8存证/交易数据时间戳不合法
    ymdNow2=$(date +'%Y-%m-%d')
    hmsNow2=$(date -d '-1 minute ago' +'%H:%M:%S')
    ymd_hms_Now2=$ymdNow2" "$hmsNow2
    globalTime123_Now2=$ymd_hms_Now2
    #时间戳不合法 没有毫秒
    dataReceipt8_1F="{\"CreateTimestamp\":\""${globalTime123_Now2}"\",\"EntityId\":\"r4252g542\",\"KeyId\":\"g52g52wgtwg\",\"ReceiptValue\":11.1,\"Version\":\"v24.0\",\"UserName\":\"sg\",\"OperationType\":\"3443g\",\"DataType\":\"3g\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"g5g542g625t2\"}"

    ymdNow3=$(date +'%Y\\%m\\%d')
    hmsNow3=$(date -d '-1 minute ago' +'%H:%M:%S')
    ymd_hms_Now3=$ymdNow3" "$hmsNow3
    globalTime123_Now3=$ymd_hms_Now3".123"
    #时间戳不合法 没有毫秒
    dataReceipt8_2F="{\"CreateTimestamp\":\""${globalTime123_Now3}"\",\"EntityId\":\"g52g52g542\",\"KeyId\":\"g5g524grwg542g\",\"ReceiptValue\":11.1,\"Version\":\"v25.0\",\"UserName\":\"sg\",\"OperationType\":\"3443g\",\"DataType\":\"3g\",\"ServiceType\":\"test\",\"FileName\":\"测试数据1.txt\",\"FileSize\":140.20,\"FileHash\":\"0xnofuegwuifgf932fou29f23992effe\",\"Uri\":\"/usr/local/1.txt\",\"ParentKeyId\":\"ey83h208ffh20f84\",\"AttachmentFileUris\":[\"e21h8e\",\"2h918eh2\"],\"AttachmentTotalHash\":\"42f432f42fhyehy\"}"

}
